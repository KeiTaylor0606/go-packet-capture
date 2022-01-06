package cap

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	snapshot_len int32         = 1024
	promiscuous  bool          = false
	timeout      time.Duration = 30 * time.Second
)

/*
GetIPAddress(): パケットキャプチャ
*/
func PacketCapture() error {
	deviceName := "ap0"                                                          // TODO: デバイス名の取得を何とかしたい
	handle, err := pcap.OpenLive(deviceName, snapshot_len, promiscuous, timeout) // デバイスの中身を開ける
	if err != nil {
		return err
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType()) // ハンドルをパケットソースとして使用し、すべてのパケットを処理する
	packetNum := 0                                                      // パケットの番号

	for packet := range packetSource.Packets() {
		if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			packetNum++ // パケット番号をインクリメント
			log.Println("")
			fmt.Println("\x1b[34mPacket Nnumber:\x1b[0m", packetNum)
			fmt.Println("\x1b[31mDestination IP:\x1b[0m", ip.DstIP)
			fmt.Println("\x1b[31mSource      IP:\x1b[0m", ip.SrcIP)
			fmt.Println("\x1b[33mCheck      Sum:\x1b[0m", ip.Checksum)
			fmt.Println("\x1b[33mLayer Contents:\x1b[0m", ip.LayerContents())
			fmt.Println("-------------------------------------")
		}
	}
	return nil
}
