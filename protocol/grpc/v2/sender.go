/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package v2

import (
	"context"
	"sync"
	"time"

	format "github.com/cloudevents/sdk-go/binding/format/protobuf/v2"
	"github.com/cloudevents/sdk-go/binding/format/protobuf/v2/pb"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/binding"
	"github.com/cloudevents/sdk-go/v2/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCTransport struct {
	client pb.CloudEventsClient
	//c      pb.CloudEventsService_SendClient
	conn *grpc.ClientConn
	m    sync.Map
}

type callback func(error)

func NewGRPCTransport(endpoint string) (*GRPCTransport, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return nil, err
	}

	t := &GRPCTransport{
		client: pb.NewCloudEventsClient(conn),
		conn:   conn,
	}
	return t, nil
}

func (t *GRPCTransport) Publish(ctx context.Context, events ...*v2.Event) error {
	var pbEvents []*pb.CloudEvent
	for idx := range events {
		e := events[idx]
		pe, err := format.ToProto(e)
		if err != nil {
			return err
		}
		pbEvents = append(pbEvents, pe)
	}

	_, err := t.client.Send(ctx, &pb.CloudEventBatch{Events: pbEvents})
	return err
}

func (t *GRPCTransport) Send(ctx context.Context, m binding.Message, transformers ...binding.Transformer) error {
	event, err := binding.ToEvent(ctx, m, transformers...)
	if err != nil {
		return err
	}
	pe, err := format.ToProto(event)
	if err != nil {
		return err
	}

	_, err = t.client.Send(ctx, &pb.CloudEventBatch{Events: []*pb.CloudEvent{pe}})
	return err
}

func (t *GRPCTransport) Request(ctx context.Context, m binding.Message, transformers ...binding.Transformer) (binding.Message, error) {
	return nil, nil
}

func (t *GRPCTransport) Respond(ctx context.Context) (binding.Message, protocol.ResponseFn, error) {
	return nil, nil, nil
}

func (t *GRPCTransport) Receive(ctx context.Context) (binding.Message, error) {
	return nil, nil
}

func (t *GRPCTransport) OpenInbound(ctx context.Context) error {
	return nil
}
