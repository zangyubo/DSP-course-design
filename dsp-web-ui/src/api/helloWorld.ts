import axios from 'axios'

// 定义提交数据的接口
export interface Post {
  name: string // 提交的用户或实体名称
}

// 定义接收数据的接口
export interface Receive {
  message: string // 返回的消息内容
}

// 提交数据并接收响应
export function sendPostRequest(data: Post): Promise<Receive> {
  return axios.post<Receive>('http://localhost:8080/hello', data).then((response) => response.data)
}
