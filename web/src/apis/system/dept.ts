import type * as System from './type'
import http from '@/utils/http'

/** @desc 获取部门数据 */
export function getSystemDeptList() {
  return http.get<System.DeptItem[]>('/system/dept')
}

/** @desc 获取部门详情 */
export function getSystemDeptDetail(params: { id: string }) {
  return http.get<System.DeptItem>('/system/dept/detail', params)
}
