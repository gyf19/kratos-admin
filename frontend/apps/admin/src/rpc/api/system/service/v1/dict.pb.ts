// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: system/service/v1/dict.proto

/* eslint-disable */
import { type Empty } from "../../../google/protobuf/empty.pb";
import { type Timestamp } from "../../../google/protobuf/timestamp.pb";
import { type PagingRequest } from "../../../pagination/v1/pagination.pb";

/** 字典 */
export interface Dict {
  /** 字典ID */
  id?:
    | number
    | null
    | undefined;
  /** 字典分类 */
  category?:
    | string
    | null
    | undefined;
  /** 字典分类名称 */
  categoryDesc?:
    | string
    | null
    | undefined;
  /** 字典键 */
  key?:
    | string
    | null
    | undefined;
  /** 字典值 */
  value?:
    | string
    | null
    | undefined;
  /** 字典值名称 */
  valueDesc?:
    | string
    | null
    | undefined;
  /** 字典值数据类型 */
  valueDataType?:
    | string
    | null
    | undefined;
  /** 字典状态 */
  status?:
    | string
    | null
    | undefined;
  /** 排序编号 */
  sortId?:
    | number
    | null
    | undefined;
  /** 备注 */
  remark?:
    | string
    | null
    | undefined;
  /** 创建时间 */
  createTime?:
    | Timestamp
    | null
    | undefined;
  /** 更新时间 */
  updateTime?:
    | Timestamp
    | null
    | undefined;
  /** 删除时间 */
  deleteTime?: Timestamp | null | undefined;
}

/** 查询字典列表 - 回应 */
export interface ListDictResponse {
  items: Dict[];
  total: number;
}

/** 查询字典详情 - 请求 */
export interface GetDictRequest {
  id: number;
}

/** 创建字典 - 请求 */
export interface CreateDictRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  data: Dict | null;
}

/** 更新字典 - 请求 */
export interface UpdateDictRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  data:
    | Dict
    | null;
  /** 要更新的字段列表 */
  updateMask:
    | string[]
    | null;
  /** 如果设置为true的时候，资源不存在则会新增(插入)，并且在这种情况下`updateMask`字段将会被忽略。 */
  allowMissing?: boolean | null | undefined;
}

/** 删除字典 - 请求 */
export interface DeleteDictRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  id: number;
}

/** 字典服务 */
export interface DictService {
  /** 查询字典列表 */
  ListDict(request: PagingRequest): Promise<ListDictResponse>;
  /** 查询字典详情 */
  GetDict(request: GetDictRequest): Promise<Dict>;
  /** 创建字典 */
  CreateDict(request: CreateDictRequest): Promise<Empty>;
  /** 更新字典 */
  UpdateDict(request: UpdateDictRequest): Promise<Empty>;
  /** 删除字典 */
  DeleteDict(request: DeleteDictRequest): Promise<Empty>;
}
