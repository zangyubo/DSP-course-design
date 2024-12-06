import axios from 'axios'
import npyjs from 'npyjs'

// 定义接收数据的接口
export interface Receive {
  message: string // 返回的消息内容
}

// 提交数据并接收响应
export function sendGetRequest(): Promise<void> {
  return axios
    .get('http://localhost:8080/testGetData', {
      responseType: 'blob', // 获取二进制数据（文件）
    })
    .then((response) => {
      // 获取返回的文件 Blob 对象
      const fileBlob = response.data

      // // 创建一个临时的 URL 用于下载文件
      // const fileUrl = URL.createObjectURL(fileBlob)

      // // 创建一个临时的下载链接并触发下载
      // const link = document.createElement('a')
      // link.href = fileUrl
      // link.download = 'time.npy' // 设置下载文件的名称
      // link.click()

      // // 释放临时的 URL 对象
      // URL.revokeObjectURL(fileUrl)

      // 读取下载的 .npy 文件并打印到控制台
      const reader = new FileReader()
      reader.onload = () => {
        const arrayBuffer = reader.result as ArrayBuffer

        // 使用 npyjs 来解析 ArrayBuffer
        const n = new npyjs()
        n.load(arrayBuffer).then((res) => {
          console.log(`🚀 ~ file: helloWorld.ts:42 ~ n.load ~ res:\n`, res)
        })
      }

      // 将 Blob 转换为 ArrayBuffer 进行处理
      reader.readAsArrayBuffer(fileBlob)
    })
}
