// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: internal_message/service/v1/notification_message.proto

/* eslint-disable */
import { type Empty } from "../../../google/protobuf/empty.pb";
import { type Timestamp } from "../../../google/protobuf/timestamp.pb";
import { type PagingRequest } from "../../../pagination/v1/pagination.pb";

/** 通知消息状态 */
export enum NotificationMessageStatus {
  /** NotificationMessageStatus_Unknown - 未知状态 */
  NotificationMessageStatus_Unknown = "NotificationMessageStatus_Unknown",
  /** NotificationMessageStatus_Draft - 草稿 */
  NotificationMessageStatus_Draft = "NotificationMessageStatus_Draft",
  /** NotificationMessageStatus_Published - 已发布 */
  NotificationMessageStatus_Published = "NotificationMessageStatus_Published",
  /** NotificationMessageStatus_Scheduled - 定时发布 */
  NotificationMessageStatus_Scheduled = "NotificationMessageStatus_Scheduled",
  /** NotificationMessageStatus_Revoked - 已撤销 */
  NotificationMessageStatus_Revoked = "NotificationMessageStatus_Revoked",
  /** NotificationMessageStatus_Archived - 已归档 */
  NotificationMessageStatus_Archived = "NotificationMessageStatus_Archived",
  /** NotificationMessageStatus_Deleted - 已删除 */
  NotificationMessageStatus_Deleted = "NotificationMessageStatus_Deleted",
}

/** 通知消息 */
export interface NotificationMessage {
  /** 消息ID */
  id?:
    | number
    | null
    | undefined;
  /** 主题 */
  subject?:
    | string
    | null
    | undefined;
  /** 内容 */
  content?:
    | string
    | null
    | undefined;
  /** 消息状态 */
  status?:
    | NotificationMessageStatus
    | null
    | undefined;
  /** 分类ID */
  categoryId?:
    | number
    | null
    | undefined;
  /** 分类名称 */
  categoryName?:
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

/** 查询通知消息列表 - 回应 */
export interface ListNotificationMessageResponse {
  items: NotificationMessage[];
  total: number;
}

/** 查询通知消息详情 - 请求 */
export interface GetNotificationMessageRequest {
  id: number;
}

/** 创建通知消息 - 请求 */
export interface CreateNotificationMessageRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  data: NotificationMessage | null;
}

/** 更新通知消息 - 请求 */
export interface UpdateNotificationMessageRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  data:
    | NotificationMessage
    | null;
  /** 要更新的字段列表 */
  updateMask:
    | string[]
    | null;
  /** 如果设置为true的时候，资源不存在则会新增(插入)，并且在这种情况下`updateMask`字段将会被忽略。 */
  allowMissing?: boolean | null | undefined;
}

/** 删除通知消息 - 请求 */
export interface DeleteNotificationMessageRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  id: number;
}

/** 通知消息服务 */
export interface NotificationMessageService {
  /** 查询通知消息列表 */
  ListNotificationMessage(request: PagingRequest): Promise<ListNotificationMessageResponse>;
  /** 查询通知消息详情 */
  GetNotificationMessage(request: GetNotificationMessageRequest): Promise<NotificationMessage>;
  /** 创建通知消息 */
  CreateNotificationMessage(request: CreateNotificationMessageRequest): Promise<Empty>;
  /** 更新通知消息 */
  UpdateNotificationMessage(request: UpdateNotificationMessageRequest): Promise<Empty>;
  /** 删除通知消息 */
  DeleteNotificationMessage(request: DeleteNotificationMessageRequest): Promise<Empty>;
}
