import * as jspb from 'google-protobuf'



export class ControlRequest extends jspb.Message {
  getInput(): number;
  setInput(value: number): ControlRequest;

  getValue(): number;
  setValue(value: number): ControlRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ControlRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ControlRequest): ControlRequest.AsObject;
  static serializeBinaryToWriter(message: ControlRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ControlRequest;
  static deserializeBinaryFromReader(message: ControlRequest, reader: jspb.BinaryReader): ControlRequest;
}

export namespace ControlRequest {
  export type AsObject = {
    input: number,
    value: number,
  }
}

export class EmptyResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EmptyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: EmptyResponse): EmptyResponse.AsObject;
  static serializeBinaryToWriter(message: EmptyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EmptyResponse;
  static deserializeBinaryFromReader(message: EmptyResponse, reader: jspb.BinaryReader): EmptyResponse;
}

export namespace EmptyResponse {
  export type AsObject = {
  }
}

