import { ref } from 'vue';

export const freX = ref([0, 1, 2, 3, 4, 5, 6, 7, 8, 9]);
export const freY = ref([10, 20, 35, 50, 70, 90, 110, 150, 180, 200]);

export const timeX = ref([]);
export const timeY = ref([]);

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
};
