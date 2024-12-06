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
    .post<UploadResponse>('http://localhost:8080/basicChartReceiveMP3', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    })
    .then((response) => response.data)
}

import npyjs from 'npyjs'

export interface send {
  channel: number
}

// 定义接收数据的接口
export interface Receive {
  message: string // 返回的消息内容
  data?: { dtype: 'float64'; data: Float64Array; shape: number[]; fortranOrder: boolean }
}

// 获取数据函数
export function getChartData(channel: number, mode: number): Promise<Receive> {
  return axios
    .get(`http://localhost:8080/basicChartSendChartData?channel=${channel}&mode=${mode}`, {
      responseType: 'blob', // 获取二进制数据（文件）
    })
    .then((response) => {
      const fileBlob = response.data

      const reader = new FileReader()
      return new Promise<Receive>((resolve) => {
        reader.onload = () => {
          const arrayBuffer = reader.result as ArrayBuffer

          // 使用 npyjs 来解析 ArrayBuffer
          const npyLoader = new npyjs()

          npyLoader
            .load(arrayBuffer)
            .then((data) => {
              resolve({
                message: 'File successfully loaded and parsed',
                data: data as never,
              })
            })
            .catch((err) => {
              console.error('Error parsing npy file:', err)
              resolve({
                message: 'Failed to parse the npy file',
              })
            })
        }

        // 将 Blob 转换为 ArrayBuffer 进行处理
        reader.readAsArrayBuffer(fileBlob)
      })
    })
    .catch((error) => {
      console.error('Error while fetching the file:', error)
      return {
        message: 'Error while fetching the file',
      }
    })
}
