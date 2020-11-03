// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package punt

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vpe "github.com/edwarnicke/govpp/pkg/v2009/binapi/vpe"
)

// RPCService defines RPC service  punt.
type RPCService interface {
	PuntReasonDump(ctx context.Context, in *PuntReasonDump) (RPCService_PuntReasonDumpClient, error)
	PuntSocketDeregister(ctx context.Context, in *PuntSocketDeregister) (*PuntSocketDeregisterReply, error)
	PuntSocketDump(ctx context.Context, in *PuntSocketDump) (RPCService_PuntSocketDumpClient, error)
	PuntSocketRegister(ctx context.Context, in *PuntSocketRegister) (*PuntSocketRegisterReply, error)
	SetPunt(ctx context.Context, in *SetPunt) (*SetPuntReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) PuntReasonDump(ctx context.Context, in *PuntReasonDump) (RPCService_PuntReasonDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_PuntReasonDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_PuntReasonDumpClient interface {
	Recv() (*PuntReasonDetails, error)
	api.Stream
}

type serviceClient_PuntReasonDumpClient struct {
	api.Stream
}

func (c *serviceClient_PuntReasonDumpClient) Recv() (*PuntReasonDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *PuntReasonDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) PuntSocketDeregister(ctx context.Context, in *PuntSocketDeregister) (*PuntSocketDeregisterReply, error) {
	out := new(PuntSocketDeregisterReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PuntSocketDump(ctx context.Context, in *PuntSocketDump) (RPCService_PuntSocketDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_PuntSocketDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_PuntSocketDumpClient interface {
	Recv() (*PuntSocketDetails, error)
	api.Stream
}

type serviceClient_PuntSocketDumpClient struct {
	api.Stream
}

func (c *serviceClient_PuntSocketDumpClient) Recv() (*PuntSocketDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *PuntSocketDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) PuntSocketRegister(ctx context.Context, in *PuntSocketRegister) (*PuntSocketRegisterReply, error) {
	out := new(PuntSocketRegisterReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SetPunt(ctx context.Context, in *SetPunt) (*SetPuntReply, error) {
	out := new(SetPuntReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}