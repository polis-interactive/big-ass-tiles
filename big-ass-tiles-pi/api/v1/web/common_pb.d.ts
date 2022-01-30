import * as jspb from 'google-protobuf'



export class EmptyRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EmptyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: EmptyRequest): EmptyRequest.AsObject;
  static serializeBinaryToWriter(message: EmptyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EmptyRequest;
  static deserializeBinaryFromReader(message: EmptyRequest, reader: jspb.BinaryReader): EmptyRequest;
}

export namespace EmptyRequest {
  export type AsObject = {
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

export class successFailure extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): successFailure;

  getError(): string;
  setError(value: string): successFailure;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): successFailure.AsObject;
  static toObject(includeInstance: boolean, msg: successFailure): successFailure.AsObject;
  static serializeBinaryToWriter(message: successFailure, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): successFailure;
  static deserializeBinaryFromReader(message: successFailure, reader: jspb.BinaryReader): successFailure;
}

export namespace successFailure {
  export type AsObject = {
    success: boolean,
    error: string,
  }
}

export class DefaultResponse extends jspb.Message {
  getStatus(): successFailure | undefined;
  setStatus(value?: successFailure): DefaultResponse;
  hasStatus(): boolean;
  clearStatus(): DefaultResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DefaultResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DefaultResponse): DefaultResponse.AsObject;
  static serializeBinaryToWriter(message: DefaultResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DefaultResponse;
  static deserializeBinaryFromReader(message: DefaultResponse, reader: jspb.BinaryReader): DefaultResponse;
}

export namespace DefaultResponse {
  export type AsObject = {
    status?: successFailure.AsObject,
  }
}

export class InteractionVibes extends jspb.Message {
  getVibetype(): VibeType;
  setVibetype(value: VibeType): InteractionVibes;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InteractionVibes.AsObject;
  static toObject(includeInstance: boolean, msg: InteractionVibes): InteractionVibes.AsObject;
  static serializeBinaryToWriter(message: InteractionVibes, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InteractionVibes;
  static deserializeBinaryFromReader(message: InteractionVibes, reader: jspb.BinaryReader): InteractionVibes;
}

export namespace InteractionVibes {
  export type AsObject = {
    vibetype: VibeType,
  }
}

export class InteractionSing extends jspb.Message {
  getOctave(): number;
  setOctave(value: number): InteractionSing;

  getNote(): number;
  setNote(value: number): InteractionSing;

  getPeak(): number;
  setPeak(value: number): InteractionSing;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InteractionSing.AsObject;
  static toObject(includeInstance: boolean, msg: InteractionSing): InteractionSing.AsObject;
  static serializeBinaryToWriter(message: InteractionSing, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InteractionSing;
  static deserializeBinaryFromReader(message: InteractionSing, reader: jspb.BinaryReader): InteractionSing;
}

export namespace InteractionSing {
  export type AsObject = {
    octave: number,
    note: number,
    peak: number,
  }
}

export class InteractionButtonMash extends jspb.Message {
  getColor(): ColorType;
  setColor(value: ColorType): InteractionButtonMash;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InteractionButtonMash.AsObject;
  static toObject(includeInstance: boolean, msg: InteractionButtonMash): InteractionButtonMash.AsObject;
  static serializeBinaryToWriter(message: InteractionButtonMash, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InteractionButtonMash;
  static deserializeBinaryFromReader(message: InteractionButtonMash, reader: jspb.BinaryReader): InteractionButtonMash;
}

export namespace InteractionButtonMash {
  export type AsObject = {
    color: ColorType,
  }
}

export class InteractionRequest extends jspb.Message {
  getType(): InteractionType;
  setType(value: InteractionType): InteractionRequest;

  getButtonmash(): InteractionButtonMash | undefined;
  setButtonmash(value?: InteractionButtonMash): InteractionRequest;
  hasButtonmash(): boolean;
  clearButtonmash(): InteractionRequest;

  getSing(): InteractionSing | undefined;
  setSing(value?: InteractionSing): InteractionRequest;
  hasSing(): boolean;
  clearSing(): InteractionRequest;

  getVibe(): InteractionVibes | undefined;
  setVibe(value?: InteractionVibes): InteractionRequest;
  hasVibe(): boolean;
  clearVibe(): InteractionRequest;

  getInteractionCase(): InteractionRequest.InteractionCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InteractionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: InteractionRequest): InteractionRequest.AsObject;
  static serializeBinaryToWriter(message: InteractionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InteractionRequest;
  static deserializeBinaryFromReader(message: InteractionRequest, reader: jspb.BinaryReader): InteractionRequest;
}

export namespace InteractionRequest {
  export type AsObject = {
    type: InteractionType,
    buttonmash?: InteractionButtonMash.AsObject,
    sing?: InteractionSing.AsObject,
    vibe?: InteractionVibes.AsObject,
  }

  export enum InteractionCase { 
    INTERACTION_NOT_SET = 0,
    BUTTONMASH = 2,
    SING = 3,
    VIBE = 4,
  }
}

export class SignState extends jspb.Message {
  getState(): SignStateType;
  setState(value: SignStateType): SignState;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SignState.AsObject;
  static toObject(includeInstance: boolean, msg: SignState): SignState.AsObject;
  static serializeBinaryToWriter(message: SignState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SignState;
  static deserializeBinaryFromReader(message: SignState, reader: jspb.BinaryReader): SignState;
}

export namespace SignState {
  export type AsObject = {
    state: SignStateType,
  }
}

export enum VibeType { 
  FIRE = 0,
  EARTH = 1,
  WATER = 2,
  AIR = 3,
}
export enum ColorType { 
  RED = 0,
  BLUE = 1,
  GREEN = 2,
  FUCHSIA = 3,
  AQUA = 4,
  YELLOW = 5,
}
export enum InteractionType { 
  NONE = 0,
  BUTTON_MASH = 1,
  SING = 2,
  VIBE = 3,
}
export enum SignStateType { 
  ON = 0,
  OFF = 1,
  DOWN = 2,
  FKD_UP = 3,
}
