import { apiClient } from './client'

export interface Memo {
  id: number
  title: string
  content: string
  pinned: boolean
  created_at: string
  updated_at: string
}

export interface CreateMemoRequest {
  title: string
  content: string
  pinned?: boolean
}

export interface UpdateMemoRequest {
  title?: string
  content?: string
  pinned?: boolean
}

export async function list(limit: number = 100, offset: number = 0): Promise<Memo[]> {
  const { data } = await apiClient.get<Memo[]>('/memos', {
    params: { limit, offset }
  })
  return data
}

export async function get(id: number): Promise<Memo> {
  const { data } = await apiClient.get<Memo>(`/memos/${id}`)
  return data
}

export async function create(input: CreateMemoRequest): Promise<Memo> {
  const { data } = await apiClient.post<Memo>('/memos', input)
  return data
}

export async function update(id: number, input: UpdateMemoRequest): Promise<Memo> {
  const { data } = await apiClient.put<Memo>(`/memos/${id}`, input)
  return data
}

export async function remove(id: number): Promise<{ message: string }> {
  const { data } = await apiClient.delete<{ message: string }>(`/memos/${id}`)
  return data
}

const memosAPI = {
  list,
  get,
  create,
  update,
  remove
}

export default memosAPI
