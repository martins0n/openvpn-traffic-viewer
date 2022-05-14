package lib

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type OpenVpnStatus struct {
	CommonName     string
	BytesReceived  int64
	BytesSent      int64
	RealAddress    string
	ConnectedSince time.Time
}

func ParseClientListRecord(record string) OpenVpnStatus {
	records := strings.Split(record, ",")
	commonName := records[0]
	realAddress := records[1]
	bytesReceived, _ := strconv.Atoi(records[2])
	bytesSent, _ := strconv.Atoi(records[3])
	connectedSince, _ := time.Parse("Mon Jan 2 15:04:05 2006", records[4])
	return OpenVpnStatus{
		CommonName:     commonName,
		BytesReceived:  int64(bytesReceived),
		BytesSent:      int64(bytesSent),
		RealAddress:    realAddress,
		ConnectedSince: connectedSince,
	}
}

func ParseOpenVpnStatus(filePath string) *[]OpenVpnStatus {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var openVpnStatusList []OpenVpnStatus

	for {
		line, err := reader.ReadString('\n')

		if strings.HasPrefix(line, "Common Name") {
			for {
				line, err := reader.ReadString('\n')
				if err != io.EOF && strings.HasPrefix(line, "ROUTING TABLE") {
					break
				} else {
					line = strings.Trim(line, "\n")
					openVpnStatusList = append(openVpnStatusList, ParseClientListRecord(line))
				}
			}
		}
		if err == io.EOF {
			break
		}
	}
	return &openVpnStatusList
}
