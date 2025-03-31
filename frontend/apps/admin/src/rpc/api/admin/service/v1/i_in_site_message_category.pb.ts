// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: admin/service/v1/i_in_site_message_category.proto

/* eslint-disable */
import { type Empty } from "../../../google/protobuf/empty.pb";
import { type PagingRequest } from "../../../pagination/v1/pagination.pb";
import {
  type CreateInSiteMessageCategoryRequest,
  type DeleteInSiteMessageCategoryRequest,
  type GetInSiteMessageCategoryRequest,
  type InSiteMessageCategory,
  type ListInSiteMessageCategoryResponse,
  type UpdateInSiteMessageCategoryRequest,
} from "../../../system/service/v1/in_site_message_category.pb";

/** 站内信分类管理服务 */
export interface InSiteMessageCategoryService {
  /** 查询站内信分类列表 */
  ListInSiteMessageCategory(request: PagingRequest): Promise<ListInSiteMessageCategoryResponse>;
  /** 查询站内信分类详情 */
  GetInSiteMessageCategory(request: GetInSiteMessageCategoryRequest): Promise<InSiteMessageCategory>;
  /** 创建站内信分类 */
  CreateInSiteMessageCategory(request: CreateInSiteMessageCategoryRequest): Promise<Empty>;
  /** 更新站内信分类 */
  UpdateInSiteMessageCategory(request: UpdateInSiteMessageCategoryRequest): Promise<Empty>;
  /** 删除站内信分类 */
  DeleteInSiteMessageCategory(request: DeleteInSiteMessageCategoryRequest): Promise<Empty>;
}
