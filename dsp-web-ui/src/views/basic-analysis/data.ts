import { ref } from 'vue'

export const xAxis1 = ref<number[]>([])
export const yAxis1 = ref<number[]>([])

export const xAxis2 = ref<number[]>([])
export const yAxis2 = ref<number[]>([])

export const exampleLable = {
  type: 'group',
  rotation: Math.PI / 4,
  bounding: 'raw',
  right: 110,
  bottom: 110,
  z: 100,
  children: [
    {
      type: 'rect',
      left: 'center',
      top: 'center',
      z: 100,
      shape: {
        width: 400,
        height: 50,
      },
      style: {
        fill: 'rgba(0,0,0,0.3)',
      },
    },
    {
      type: 'text',
      left: 'center',
      top: 'center',
      z: 100,
      style: {
        fill: '#fff',
        text: '示例图表',
        font: 'bold 26px sans-serif',
      },
    },
  ],
}
