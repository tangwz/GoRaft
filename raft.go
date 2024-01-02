package goraft

import (
	"fmt"
	"log"
	"sync"
)

const DebugRaft = true

type Raft struct {
	mu sync.Mutex

	// id is the server ID of this Raft
	id int

	// peerIds is the list of server IDs of all peers
	peerIds []int

	// server is the server containing this Raft. It's used to send RPC calls to peers
	server *Server
}

func NewRaft(id int, peerIds []int, server *Server, applyCh <-chan interface{}) *Raft {
	rf := new(Raft)
	rf.id = id
	rf.peerIds = peerIds
	rf.server = server

	return rf
}

func (rf *Raft) Stop() {
}

// dlog logs a debugging message if DebugRaft is true
func (rf *Raft) dlog(format string, args ...interface{}) {
	if DebugRaft {
		format = fmt.Sprintf("[%d] ", rf.id) + format
		log.Printf(format, args...)
	}
}

type RequestVoteArgs struct {
}

type RequestVoteReply struct {
}

func (rf *Raft) RequestVote(args RequestVoteArgs, reply *RequestVoteReply) error {
	return nil
}

type AppendEntriesArgs struct {
}

type AppendEntriesReply struct {
}

func (rf *Raft) AppendEntries(args AppendEntriesArgs, reply *AppendEntriesReply) error {
	return nil
}
