// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: admin/service/v1/i_notification_message_category.proto

/* eslint-disable */
import { type Empty } from "../../../google/protobuf/empty.pb";
import {
  type CreateNotificationMessageCategoryRequest,
  type DeleteNotificationMessageCategoryRequest,
  type GetNotificationMessageCategoryRequest,
  type ListNotificationMessageCategoryResponse,
  type NotificationMessageCategory,
  type UpdateNotificationMessageCategoryRequest,
} from "../../../internal_message/service/v1/notification_message_category.pb";
import { type PagingRequest } from "../../../pagination/v1/pagination.pb";

/** 通知消息分类管理服务 */
export interface NotificationMessageCategoryService {
  /** 查询通知消息分类列表 */
  ListNotificationMessageCategory(request: PagingRequest): Promise<ListNotificationMessageCategoryResponse>;
  /** 查询通知消息分类详情 */
  GetNotificationMessageCategory(request: GetNotificationMessageCategoryRequest): Promise<NotificationMessageCategory>;
  /** 创建通知消息分类 */
  CreateNotificationMessageCategory(request: CreateNotificationMessageCategoryRequest): Promise<Empty>;
  /** 更新通知消息分类 */
  UpdateNotificationMessageCategory(request: UpdateNotificationMessageCategoryRequest): Promise<Empty>;
  /** 删除通知消息分类 */
  DeleteNotificationMessageCategory(request: DeleteNotificationMessageCategoryRequest): Promise<Empty>;
}
