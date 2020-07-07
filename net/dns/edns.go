package dns

import (
	"bytes"
	"net"
)

type EDNSOPT [2]byte

//https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-11
var (
	Reserved         = EDNSOPT{0b00000000, 0b00000000} //0
	LLQ              = EDNSOPT{0b00000000, 0b00000001} //1 Optional
	UL               = EDNSOPT{0b00000000, 0b00000010} //2 On-hold
	NSID             = EDNSOPT{0b00000000, 0b00000011} //3 Standard
	Reserved2        = EDNSOPT{0b00000000, 0b00000100} //4
	DAU              = EDNSOPT{0b00000000, 0b00000101} //5 Standard
	DHU              = EDNSOPT{0b00000000, 0b00000110} //6 Standard
	N3U              = EDNSOPT{0b00000000, 0b00000111} //7 Standard
	EdnsClientSubnet = EDNSOPT{0b00000000, 0b00001000} //8 Optional
	EDNSEXPIRE       = EDNSOPT{0b00000000, 0b00001001} //9 Optional
	COOKIE           = EDNSOPT{0b00000000, 0b00001010} //10 Standard
	TcpKeepalive     = EDNSOPT{0b00000000, 0b00001011} //11 Standard
	Padding          = EDNSOPT{0b00000000, 0b00001100} //12 Standard
	CHAIN            = EDNSOPT{0b00000000, 0b00001101} //13 Standard
	KEYTAG           = EDNSOPT{0b00000000, 0b00001110} //14 Optional
	ExtendedDNSError = EDNSOPT{0b00000000, 0b00001111} //15 Standard
	EDNSClientTag    = EDNSOPT{0b00000000, 0b00010000} //16 Optional
	EDNSServerTag    = EDNSOPT{0b00000000, 0b00010001} //17 Optional

)

type eDNSHeader struct {
	DnsHeader   []byte
	Name        [1]byte
	Type        [2]byte
	PayloadSize [2]byte
	ExtendRCode [1]byte
	EDNSVersion [1]byte
	Z           [2]byte
	Data        []byte
}

func createEDNSRequ(header eDNSHeader) []byte {
	header.DnsHeader[10] = 0b00000000
	header.DnsHeader[11] = 0b00000001
	//log.Println(header.DnsHeader)
	length := getLength(len(header.Data))
	return bytes.Join([][]byte{header.DnsHeader, header.Name[:], header.Type[:], header.PayloadSize[:], header.ExtendRCode[:], header.EDNSVersion[:], header.Z[:], length[:], header.Data}, []byte{})
}

// https://tools.ietf.org/html/rfc7871
// https://tools.ietf.org/html/rfc2671
func createEDNSReq(domain string, reqType2 reqType, eDNS []byte) []byte {
	data := bytes.NewBuffer(creatRequest(domain, reqType2, true))
	data.WriteByte(0b00000000)
	data.Write([]byte{0b00000000, 0b00101001})
	data.Write([]byte{0b00010000, 0b00000000})
	data.WriteByte(0b00000000)
	data.WriteByte(0b00000000)
	data.Write([]byte{0b00000000, 0b00000000})
	//name := []byte{0b00000000}
	//typeR := []byte{0b00000000, 0b00101001}       //41
	//payloadSize := []byte{0b00010000, 0b00000000} //4096
	//extendRcode := []byte{0b00000000}
	//eDNSVersion := []byte{0b00000000}
	//z := []byte{0b00000000, 0b00000000}
	data.Write(getLength(len(eDNS)))
	data.Write(eDNS)
	//return bytes.Join([][]byte{normalReq, name, typeR, payloadSize, extendRcode, eDNSVersion, z, dataLength[:], eDNS}, []byte{})
	return data.Bytes()
}

func createEdnsClientSubnet(ip *net.IPNet) []byte {
	optionData := bytes.NewBuffer(nil)
	mask, _ := ip.Mask.Size()
	//family := []byte{0b00000000, 0b00000001} // 1:Ipv4 2:IPv6 https://www.iana.org/assignments/address-family-numbers/address-family-numbers.xhtml
	//sourceNetmask := []byte{byte(mask)}      // 32
	//scopeNetmask := []byte{0b00000000}       //0 In queries, it MUST be set to 0.
	subnet := ip.IP.To4()
	if subnet == nil {
		optionData.Write([]byte{0b00000000, 0b00000010}) // family ipv6 2
		subnet = ip.IP.To16()
	} else {
		optionData.Write([]byte{0b00000000, 0b00000001}) // family ipv4 1
	}
	optionData.WriteByte(byte(mask)) // mask
	optionData.WriteByte(0b00000000) // 0 In queries, it MUST be set to 0.
	optionData.Write(subnet)         // subnet
	//optionData := bytes.Join([][]byte{family, sourceNetmask, scopeNetmask, subnet}, []byte{})

	//optionCode := []byte{EdnsClientSubnet[0], EdnsClientSubnet[1]}
	//optionLength := getLength(len(optionData))
	data := bytes.NewBuffer(nil)
	data.Write([]byte{EdnsClientSubnet[0], EdnsClientSubnet[1]}) // option Code
	data.Write(getLength(optionData.Len()))                      // option data length
	data.Write(optionData.Bytes())                               // option data
	//return bytes.Join([][]byte{optionCode, optionLength, optionData}, []byte{})
	return data.Bytes()
}

func getLength(length int) []byte {
	return []byte{byte(length >> 8), byte(length - ((length >> 8) << 8))}
}

func resolveAdditional(b []byte, arCount int) {
	for arCount != 0 {
		arCount--
		//name := b[:1]
		b = b[1:]
		typeE := b[:2]
		b = b[2:]
		//payLoadSize := b[:2]
		b = b[2:]
		//rCode := b[:1]
		b = b[1:]
		//version := b[:1]
		b = b[1:]
		//z := b[:2]
		b = b[2:]
		dataLength := int(b[0])<<8 + int(b[1])
		b = b[2:]
		//log.Println(name, typeE, payLoadSize, rCode, version, z, dataLength)
		if typeE[0] != 0 || typeE[1] != 41 {
			//optData := b[:dataLength]
			b = b[dataLength:]
			continue
		}

		if dataLength == 0 {
			continue
		}
		optCode := EDNSOPT{b[0], b[1]}
		b = b[2:]
		optionLength := int(b[0])<<8 + int(b[1])
		b = b[2:]
		switch optCode {
		case EdnsClientSubnet:
			//family := b[:2]
			b = b[2:]
			//sourceNetmask := b[:1]
			//log.Println("sourceNetmask", sourceNetmask)
			b = b[1:]
			//scopeNetmask := b[:1]
			//log.Println("scopeNetmask", scopeNetmask)
			b = b[1:]
			// Subnet IP
			//if family[0] == 0 && family[1] == 1 {
			//	log.Println(b[:4])
			//}
			//if family[0] == 0 && family[1] == 2 {
			//	log.Println(b[:16])
			//}

			b = b[optionLength-4:]
		default:
			//log.Println("opt data:", b[:optionLength])
			b = b[optionLength:]
		}
	}

}
