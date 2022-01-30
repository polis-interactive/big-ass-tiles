/**
 * @fileoverview gRPC-Web generated client stub for SignageBackend.v1.participant
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as common_pb from './common_pb';
import * as participant_pb from './participant_pb';


export class ParticipantClient {
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

  methodInfoRequestInteraction = new grpcWeb.AbstractClientBase.MethodInfo(
    common_pb.DefaultResponse,
    (request: participant_pb.InteractionRequest) => {
      return request.serializeBinary();
    },
    common_pb.DefaultResponse.deserializeBinary
  );

  requestInteraction(
    request: participant_pb.InteractionRequest,
    metadata: grpcWeb.Metadata | null): Promise<common_pb.DefaultResponse>;

  requestInteraction(
    request: participant_pb.InteractionRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: common_pb.DefaultResponse) => void): grpcWeb.ClientReadableStream<common_pb.DefaultResponse>;

  requestInteraction(
    request: participant_pb.InteractionRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: common_pb.DefaultResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/SignageBackend.v1.participant.Participant/RequestInteraction',
        request,
        metadata || {},
        this.methodInfoRequestInteraction,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/SignageBackend.v1.participant.Participant/RequestInteraction',
    request,
    metadata || {},
    this.methodInfoRequestInteraction);
  }

  methodInfoParticipantConnection = new grpcWeb.AbstractClientBase.MethodInfo(
    participant_pb.ConnectionResponse,
    (request: common_pb.EmptyRequest) => {
      return request.serializeBinary();
    },
    participant_pb.ConnectionResponse.deserializeBinary
  );

  participantConnection(
    request: common_pb.EmptyRequest,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/SignageBackend.v1.participant.Participant/ParticipantConnection',
      request,
      metadata || {},
      this.methodInfoParticipantConnection);
  }

}

