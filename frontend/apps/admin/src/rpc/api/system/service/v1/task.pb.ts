// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: system/service/v1/task.proto

/* eslint-disable */
import { type Duration } from "../../../google/protobuf/duration.pb";
import { type Empty } from "../../../google/protobuf/empty.pb";
import { type Timestamp } from "../../../google/protobuf/timestamp.pb";
import { type PagingRequest } from "../../../pagination/v1/pagination.pb";

/** 调度任务类型 */
export enum TaskType {
  /** TaskType_Periodic - 周期性任务 */
  TaskType_Periodic = "TaskType_Periodic",
  /** TaskType_Delay - 延时任务 */
  TaskType_Delay = "TaskType_Delay",
  /** TaskType_WaitResult - 等待结果 */
  TaskType_WaitResult = "TaskType_WaitResult",
}

/** 调度任务控制类型 */
export enum TaskControlType {
  /** ControlType_Start - 启动 */
  ControlType_Start = "ControlType_Start",
  /** ControlType_Stop - 停止 */
  ControlType_Stop = "ControlType_Stop",
  /** ControlType_Restart - 重启 */
  ControlType_Restart = "ControlType_Restart",
}

/** 调度任务 */
export interface Task {
  /** 任务ID */
  id?:
    | number
    | null
    | undefined;
  /** 任务类型 */
  type?:
    | TaskType
    | null
    | undefined;
  /** 任务执行类型名 */
  typeName?:
    | string
    | null
    | undefined;
  /** 任务的参数，以 JSON 格式存储，方便存储不同类型和数量的参数 */
  taskPayload?:
    | string
    | null
    | undefined;
  /** cron表达式 */
  cronSpec?:
    | string
    | null
    | undefined;
  /** 任务最多可以重试的次数 */
  retryCount?:
    | number
    | null
    | undefined;
  /** 任务超时时间 */
  timeout?:
    | Duration
    | null
    | undefined;
  /** 任务截止时间 */
  deadline?:
    | Timestamp
    | null
    | undefined;
  /** 任务延迟处理时间 */
  processIn?:
    | Duration
    | null
    | undefined;
  /** 任务执行时间点 */
  processAt?:
    | Timestamp
    | null
    | undefined;
  /** 启用/禁用任务 */
  enable?:
    | boolean
    | null
    | undefined;
  /** 备注 */
  remark?:
    | string
    | null
    | undefined;
  /** 创建者ID */
  createBy?:
    | number
    | null
    | undefined;
  /** 更新者ID */
  updateBy?:
    | number
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

/** 查询调度任务列表 - 回应 */
export interface ListTaskResponse {
  items: Task[];
  total: number;
}

/** 查询调度任务详情 - 请求 */
export interface GetTaskRequest {
  id: number;
}

export interface GetTaskByTypeNameRequest {
  /** 任务执行类型名 */
  typeName: string;
}

/** 创建调度任务 - 请求 */
export interface CreateTaskRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  data: Task | null;
}

/** 更新调度任务 - 请求 */
export interface UpdateTaskRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  data:
    | Task
    | null;
  /** 要更新的字段列表 */
  updateMask:
    | string[]
    | null;
  /** 如果设置为true的时候，资源不存在则会新增(插入)，并且在这种情况下`updateMask`字段将会被忽略。 */
  allowMissing?: boolean | null | undefined;
}

/** 删除调度任务 - 请求 */
export interface DeleteTaskRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  id: number;
}

/** 重启调度任务 - 回应 */
export interface RestartAllTaskResponse {
  count: number;
}

/** 控制调度任务 - 请求 */
export interface ControlTaskRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  controlType: TaskControlType;
  /** 任务执行类型名 */
  typeName: string;
}

/** 调度任务服务 */
export interface TaskService {
  /** 查询调度任务列表 */
  ListTask(request: PagingRequest): Promise<ListTaskResponse>;
  /** 查询调度任务详情 */
  GetTask(request: GetTaskRequest): Promise<Task>;
  GetTaskByTypeName(request: GetTaskByTypeNameRequest): Promise<Task>;
  /** 创建调度任务 */
  CreateTask(request: CreateTaskRequest): Promise<Empty>;
  /** 更新调度任务 */
  UpdateTask(request: UpdateTaskRequest): Promise<Empty>;
  /** 删除调度任务 */
  DeleteTask(request: DeleteTaskRequest): Promise<Empty>;
  /** 重启所有的调度任务 */
  RestartAllTask(request: Empty): Promise<RestartAllTaskResponse>;
  /** 停止所有的调度任务 */
  StopAllTask(request: Empty): Promise<Empty>;
  /** 控制调度任务 */
  ControlTask(request: ControlTaskRequest): Promise<Empty>;
}
