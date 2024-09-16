package main

import (
	"crypto/rand"
	"encoding/binary"
	ran "math/rand"
	"os"

	"github.com/jackpal/bencode-go"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type torrentInfo struct {
	Announce  string
	info_hash string
}

// func udpSend() {

// }
// let a: int = 5
var id = make([]byte, 20)
var madeId = 0

func genID() []byte {
	if madeId == 0 {
		rand.Read(id[8:20])
		// id = ([]byte("-AT0001-"))
		madeId = 1
	}
	return id
}

func buildConnRequest() []byte {
	// provides the string to send to server for initial connection

	// Offset  Size            Name            Value
	// 0       64-bit integer  connection_id   0x41727101980
	// 8       32-bit integer  action          0 (0 means connect)
	// 12      32-bit integer  transaction_id  ? // random

	// <Buffer 00 00 04 17 27 10 19 80 00 00 00 00 a6 ec 6b 7d>

	buf := make([]byte, 16)
	buf[2] = 0x4
	buf[3] = 0x17
	buf[4] = 0x27
	buf[5] = 0x10
	buf[6] = 0x19
	buf[7] = 0x80
	// TODO last 4 bytes should be random
	buf[12] = 0xa6
	buf[13] = 0xec
	buf[14] = 0x6b
	buf[15] = 0x7d

	return buf

}

func parseConnResponse(response []byte) (uint32, uint32, uint64) {
	// function parseConnResp(resp) {
	// return {
	//   action: resp.readUInt32BE(0),
	//   transactionId: resp.readUInt32BE(4),
	//   connectionId: resp.slice(8)
	// }
	//   }

	action := binary.BigEndian.Uint32(response[0:4])
	transactionId := binary.BigEndian.Uint32(response[4:8])
	connectionId := binary.BigEndian.Uint64(response[8:16])

	return action, transactionId, connectionId

}

func buildAnnounceRequest(connectionId uint64, info_hash []byte,
	peer_id []byte, port uint16,
) []byte {
	// Offset  Size    Name    Value
	// 0       64-bit integer  connection_id
	// 8       32-bit integer  action          1 // announce
	// 12      32-bit integer  transaction_id
	// 16      20-byte string  info_hash
	// 36      20-byte string  peer_id
	// 56      64-bit integer  downloaded
	// 64      64-bit integer  left
	// 72      64-bit integer  uploaded
	// 80      32-bit integer  event           0 // 0: none; 1: completed; 2: started; 3: stopped
	// 84      32-bit integer  IP address      0 // default
	// 88      32-bit integer  key             ? // random
	// 92      32-bit integer  num_want        -1 // default
	// 96      16-bit integer  port            ? // should be betwee
	// 98

	buffer := make([]byte, 0)

	buffer = binary.BigEndian.AppendUint64(buffer, connectionId)
	buffer = binary.BigEndian.AppendUint32(buffer, 1)

	transactionId := ran.Uint32()
	buffer = binary.BigEndian.AppendUint32(buffer, transactionId)

	buffer = append(buffer, info_hash...)
	buffer = append(buffer, peer_id...)
	buffer = binary.BigEndian.AppendUint64(buffer, 0)

	// left: buffer = binary.

	buffer = binary.BigEndian.AppendUint64(buffer, 0)
	buffer = binary.BigEndian.AppendUint32(buffer, 0)
	buffer = binary.BigEndian.AppendUint32(buffer, 0)

	key := ran.Uint32()
	buffer = binary.BigEndian.AppendUint32(buffer, key)

	//handle negative value buffer = binary.BigEndian.AppendUint32(buffer, -1)

	buffer = binary.BigEndian.AppendUint16(buffer, port)

	if len(buffer) != 98 {
		panic("s")
	}

	return buffer
	// fmt.Printf("%v", buffer)
	// println("\n")
	// print(len(buffer))

}

func main() {

	// fmt.Printf("%v", genID())

	file, err := os.Open("puppy.torrent")
	check(err)

	var torrent = torrentInfo{"", ""}
	bencode.Unmarshal(file, &torrent)

	// code for raw printing torrent file
	// file_contents, err := os.ReadFile("puppy.torrent")
	// check(err)

	// fmt.Printf("%v", string(file_contents))

	// torrentURL, err := url.Parse(torrent.Announce)
	// check(err)

	// torrentURLIP, err := net.ResolveUDPAddr("udp", torrentURL.Host)
	// check(err)

	// fmt.Printf("%s", torrentURLIP.String())

	// conn, err := net.Dial("udp", torrentURLIP.String())
	// check(err)

	// // msg := "string"

	// // fmt.Fprintf(conn, msg)
	// // (either?)
	// conn.Write(buildConnRequest())

	// readBuffer := make([]byte, 16)

	// _, err = bufio.NewReader(conn).Read(readBuffer)
	// check(err)

	// parseConnResponse(readBuffer)

	// buildAnnounceRequest(12)

}
