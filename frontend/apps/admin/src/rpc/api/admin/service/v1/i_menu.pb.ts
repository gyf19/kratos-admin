// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: admin/service/v1/i_menu.proto

/* eslint-disable */
import { type Empty } from "../../../google/protobuf/empty.pb";
import { type PagingRequest } from "../../../pagination/v1/pagination.pb";
import {
  type CreateMenuRequest,
  type DeleteMenuRequest,
  type GetMenuRequest,
  type ListMenuResponse,
  type Menu,
  type UpdateMenuRequest,
} from "../../../system/service/v1/menu.pb";

/** 后台菜单管理服务 */
export interface MenuService {
  /** 查询菜单列表 */
  ListMenu(request: PagingRequest): Promise<ListMenuResponse>;
  /** 查询菜单详情 */
  GetMenu(request: GetMenuRequest): Promise<Menu>;
  /** 创建菜单 */
  CreateMenu(request: CreateMenuRequest): Promise<Empty>;
  /** 更新菜单 */
  UpdateMenu(request: UpdateMenuRequest): Promise<Empty>;
  /** 删除菜单 */
  DeleteMenu(request: DeleteMenuRequest): Promise<Empty>;
}
