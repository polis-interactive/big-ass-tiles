import * as jspb from 'google-protobuf'

import * as common_pb from './common_pb';


export class InteractionRequest extends jspb.Message {
  getParticipantid(): number;
  setParticipantid(value: number): InteractionRequest;

  getRequest(): common_pb.InteractionRequest | undefined;
  setRequest(value?: common_pb.InteractionRequest): InteractionRequest;
  hasRequest(): boolean;
  clearRequest(): InteractionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InteractionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: InteractionRequest): InteractionRequest.AsObject;
  static serializeBinaryToWriter(message: InteractionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InteractionRequest;
  static deserializeBinaryFromReader(message: InteractionRequest, reader: jspb.BinaryReader): InteractionRequest;
}

export namespace InteractionRequest {
  export type AsObject = {
    participantid: number,
    request?: common_pb.InteractionRequest.AsObject,
  }
}

export class ConnectionAssignment extends jspb.Message {
  getConnectionid(): number;
  setConnectionid(value: number): ConnectionAssignment;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConnectionAssignment.AsObject;
  static toObject(includeInstance: boolean, msg: ConnectionAssignment): ConnectionAssignment.AsObject;
  static serializeBinaryToWriter(message: ConnectionAssignment, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConnectionAssignment;
  static deserializeBinaryFromReader(message: ConnectionAssignment, reader: jspb.BinaryReader): ConnectionAssignment;
}

export namespace ConnectionAssignment {
  export type AsObject = {
    connectionid: number,
  }
}

export class QueuePosition extends jspb.Message {
  getConnectionposition(): number;
  setConnectionposition(value: number): QueuePosition;

  getQueuelength(): number;
  setQueuelength(value: number): QueuePosition;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QueuePosition.AsObject;
  static toObject(includeInstance: boolean, msg: QueuePosition): QueuePosition.AsObject;
  static serializeBinaryToWriter(message: QueuePosition, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QueuePosition;
  static deserializeBinaryFromReader(message: QueuePosition, reader: jspb.BinaryReader): QueuePosition;
}

export namespace QueuePosition {
  export type AsObject = {
    connectionposition: number,
    queuelength: number,
  }
}

export class ConnectionResponse extends jspb.Message {
  getConnectionassignment(): ConnectionAssignment | undefined;
  setConnectionassignment(value?: ConnectionAssignment): ConnectionResponse;
  hasConnectionassignment(): boolean;
  clearConnectionassignment(): ConnectionResponse;

  getSignstate(): common_pb.SignState | undefined;
  setSignstate(value?: common_pb.SignState): ConnectionResponse;
  hasSignstate(): boolean;
  clearSignstate(): ConnectionResponse;

  getQueueposition(): QueuePosition | undefined;
  setQueueposition(value?: QueuePosition): ConnectionResponse;
  hasQueueposition(): boolean;
  clearQueueposition(): ConnectionResponse;

  getIncontrol(): common_pb.EmptyRequest | undefined;
  setIncontrol(value?: common_pb.EmptyRequest): ConnectionResponse;
  hasIncontrol(): boolean;
  clearIncontrol(): ConnectionResponse;

  getResponseCase(): ConnectionResponse.ResponseCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConnectionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ConnectionResponse): ConnectionResponse.AsObject;
  static serializeBinaryToWriter(message: ConnectionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConnectionResponse;
  static deserializeBinaryFromReader(message: ConnectionResponse, reader: jspb.BinaryReader): ConnectionResponse;
}

export namespace ConnectionResponse {
  export type AsObject = {
    connectionassignment?: ConnectionAssignment.AsObject,
    signstate?: common_pb.SignState.AsObject,
    queueposition?: QueuePosition.AsObject,
    incontrol?: common_pb.EmptyRequest.AsObject,
  }

  export enum ResponseCase { 
    RESPONSE_NOT_SET = 0,
    CONNECTIONASSIGNMENT = 2,
    SIGNSTATE = 3,
    QUEUEPOSITION = 4,
    INCONTROL = 5,
  }
}

