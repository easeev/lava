/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { GenesisState as GenesisState1 } from "../common/fixationEntry";
import { Params } from "./params";

export const protobufPackage = "lavanet.lava.projects";

/** GenesisState defines the projects module's genesis state. */
export interface GenesisState {
  params?: Params;
  projectsFS?: GenesisState1;
  /** this line is used by starport scaffolding # genesis/proto/state */
  developerFS?: GenesisState1;
}

function createBaseGenesisState(): GenesisState {
  return { params: undefined, projectsFS: undefined, developerFS: undefined };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.projectsFS !== undefined) {
      GenesisState1.encode(message.projectsFS, writer.uint32(18).fork()).ldelim();
    }
    if (message.developerFS !== undefined) {
      GenesisState1.encode(message.developerFS, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag != 10) {
            break;
          }

          message.params = Params.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag != 18) {
            break;
          }

          message.projectsFS = GenesisState1.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag != 26) {
            break;
          }

          message.developerFS = GenesisState1.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) == 4 || tag == 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      projectsFS: isSet(object.projectsFS) ? GenesisState1.fromJSON(object.projectsFS) : undefined,
      developerFS: isSet(object.developerFS) ? GenesisState1.fromJSON(object.developerFS) : undefined,
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.projectsFS !== undefined &&
      (obj.projectsFS = message.projectsFS ? GenesisState1.toJSON(message.projectsFS) : undefined);
    message.developerFS !== undefined &&
      (obj.developerFS = message.developerFS ? GenesisState1.toJSON(message.developerFS) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GenesisState>, I>>(base?: I): GenesisState {
    return GenesisState.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.projectsFS = (object.projectsFS !== undefined && object.projectsFS !== null)
      ? GenesisState1.fromPartial(object.projectsFS)
      : undefined;
    message.developerFS = (object.developerFS !== undefined && object.developerFS !== null)
      ? GenesisState1.fromPartial(object.developerFS)
      : undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends Array<infer U> ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
