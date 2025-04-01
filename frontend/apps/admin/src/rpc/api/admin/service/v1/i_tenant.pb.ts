// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: admin/service/v1/i_tenant.proto

/* eslint-disable */
import { type Empty } from "../../../google/protobuf/empty.pb";
import { type PagingRequest } from "../../../pagination/v1/pagination.pb";
import {
  type CreateTenantRequest,
  type DeleteTenantRequest,
  type GetTenantRequest,
  type ListTenantResponse,
  type Tenant,
  type UpdateTenantRequest,
} from "../../../user/service/v1/tenant.pb";

/** 租户管理服务 */
export interface TenantService {
  /** 获取租户列表 */
  ListTenant(request: PagingRequest): Promise<ListTenantResponse>;
  /** 获取租户数据 */
  GetTenant(request: GetTenantRequest): Promise<Tenant>;
  /** 创建租户 */
  CreateTenant(request: CreateTenantRequest): Promise<Empty>;
  /** 更新租户 */
  UpdateTenant(request: UpdateTenantRequest): Promise<Empty>;
  /** 删除租户 */
  DeleteTenant(request: DeleteTenantRequest): Promise<Empty>;
}
