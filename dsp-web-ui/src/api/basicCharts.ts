import axios from 'axios'

// 定义上传文件的接口
export interface UploadFile {
  file: File // 要上传的文件
  channel: number
}

// 定义接收响应数据的接口
export interface UploadResponse {
  message: string // 返回的消息
}

// 提交文件上传请求并接收响应
export function uploadFileRequest(data: UploadFile): Promise<UploadResponse> {
  const formData = new FormData()
  formData.append('file', data.file)
  formData.append('channel', data.channel.toString())

  return axios
    .post<UploadResponse>('http://localhost:8080/basicCharts', formData, {
      headers: {
        'Content-Type': 'multipart/form-data', // 设置请求头为文件上传
      },
    })
    .then((response) => response.data)
}
