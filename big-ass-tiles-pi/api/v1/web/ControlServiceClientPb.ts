/**
 * @fileoverview gRPC-Web generated client stub for BigAssTilesPi.v1.control
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as control_pb from './control_pb';


export class ControlClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoRequestControl = new grpcWeb.AbstractClientBase.MethodInfo(
    control_pb.EmptyResponse,
    (request: control_pb.ControlRequest) => {
      return request.serializeBinary();
    },
    control_pb.EmptyResponse.deserializeBinary
  );

  requestControl(
    request: control_pb.ControlRequest,
    metadata: grpcWeb.Metadata | null): Promise<control_pb.EmptyResponse>;

  requestControl(
    request: control_pb.ControlRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: control_pb.EmptyResponse) => void): grpcWeb.ClientReadableStream<control_pb.EmptyResponse>;

  requestControl(
    request: control_pb.ControlRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: control_pb.EmptyResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/BigAssTilesPi.v1.control.Control/RequestControl',
        request,
        metadata || {},
        this.methodInfoRequestControl,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/BigAssTilesPi.v1.control.Control/RequestControl',
    request,
    metadata || {},
    this.methodInfoRequestControl);
  }

}

