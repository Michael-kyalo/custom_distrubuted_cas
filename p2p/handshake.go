package p2p

type ShakeHandsFunc func(Peer) error

func NOShakeNeeded(Peer) error { return nil }
