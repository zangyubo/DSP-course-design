<template>
  <div
    style="
      width: 100%;
      height: 90%;
      background-color: white;
      display: flex;
      align-items: center;
      justify-content: center;
      flex-direction: column;
    "
  >
    <div
      style="
        height: 40%;
        width: 99%;
        margin: 3px 0px;
        display: flex;
        flex-direction: row;
      "
    >
      <div
        style="
          width: 50%;
          height: 100%;
          display: flex;
          align-items: center;
          justify-content: left;
        "
      >
        <a-button type="primary" style="margin: 0px 3px; height: 32px">
          原始音频上传
        </a-button>

        <a-button type="primary" style="margin: 0px 3px; height: 32px">
          分析信号
        </a-button>
        <audio
          style="height: 32px; width: 320px; margin: 0px 3px"
          controls
        ></audio>
        <a-select
          v-model="modeValue"
          style="height: 32px; width: 150px; margin: 0px 2px"
          placeholder="选择噪声模式"
          value-key="key"
          @change="modeChange"
        >
          <a-option
            v-for="item of selectedOption"
            :key="item.key"
            :value="item"
            :label="item.label"
          />
        </a-select>
      </div>
      <div
        style="
          width: 50%;
          height: 100%;
          display: flex;
          align-items: center;
          justify-content: center;
        "
      >
        <!-- sin noise part -->
        <a-row :gutter="16" :style="{ display: isSinNoise ? 'flex' : 'none' }">
          <a-col :span="6">
            <a-input-number v-model="aValue" style="height: 32px">
              <template #prefix> A</template></a-input-number
            >
          </a-col>

          <a-col :span="6">
            <a-input-number v-model="fValue" style="height: 32px">
              <template #prefix> f</template></a-input-number
            >
          </a-col>
          <a-col :span="6">
            <a-input-number v-model="phiValue" style="height: 32px">
              <template #prefix><p>&phi;</p></template></a-input-number
            >
          </a-col>
          <a-col :span="6">
            <a-button
              style="height: 32px; width: 100%"
              type="primary"
              :disabled="disableButton"
              @click="addSinNoise"
            >
              {{ disableButton ? '' : '添加正弦噪声' }}
            </a-button>
          </a-col>
        </a-row>

        <!-- white noise part -->
        <a-row :gutter="16" :style="{ display: isSinNoise ? 'none' : 'flex' }">
          <a-col :span="6">
            <a-input-number v-model="sValue" style="height: 32px">
              <template #prefix> S</template>
            </a-input-number>
          </a-col>

          <a-col :span="6">
            <a-input-number v-model="eValue" style="height: 32px">
              <template #prefix> E</template></a-input-number
            >
          </a-col>
          <a-col :span="6">
            <a-input-number style="height: 32px" :disabled="true">
              <template #prefix> </template
            ></a-input-number>
          </a-col>
          <a-col :span="6">
            <a-button
              style="height: 32px; width: 100%"
              type="primary"
              @click="addWhiteNoise"
            >
              添加白噪声
            </a-button>
          </a-col>
        </a-row>
      </div>
    </div>

    <div class="function-bottom">
      <a-button style="margin: 0px 3px" type="primary"> 生成噪声信号 </a-button>

      <a-input-tag
        v-model="noiseArray"
        style="height: 32px; margin: 0px 3px"
        placeholder="Please Enter"
        @remove="removeNoiseItem"
      />
      <audio
        style="height: 32px; width: 600px; margin: 0px 3px"
        controls
      ></audio>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { Message } from '@arco-design/web-vue';
  import { ref } from 'vue';

  // sin noise part var
  const aValue = ref<number>();
  const fValue = ref<number>();
  const phiValue = ref<number>();

  // white noise part var
  const sValue = ref<number>();
  const eValue = ref<number>();

  const isSinNoise = ref(true);
  const disableButton = ref(true);

  const modeValue = ref();
  const selectedOption = [
    { values: '正弦噪声', label: '正弦噪声', key: '1' },
    { values: '白噪声', label: '白噪声', key: '2' },
  ];
  const noiseArray = ref<string[]>([]);

  const inputBoxInit = () => {
    noiseArray.value = [];
    aValue.value = undefined;
    fValue.value = undefined;
    phiValue.value = undefined;
    sValue.value = undefined;
    eValue.value = undefined;
  };

  const modeChange = () => {
    disableButton.value = false;

    inputBoxInit();

    if (modeValue.value.key === '1') {
      isSinNoise.value = true;
    } else {
      isSinNoise.value = false;
    }
  };

  const addSinNoise = () => {
    if (!aValue.value) {
      Message.error('请输入幅值');
      return;
    }

    if (!fValue.value) {
      Message.error('请输入频率');
      return;
    }

    if (!phiValue.value) {
      Message.error('请输入相位');
      return;
    }

    const noiseSignal = `${aValue.value.toString()}sin(2π · (${fValue.value.toString()} · t + ${phiValue.value.toString()}))`;

    noiseArray.value.push(noiseSignal);

    aValue.value = undefined;
    fValue.value = undefined;
    phiValue.value = undefined;
  };

  const removeNoiseItem = (removedTag: string | number) => {
    noiseArray.value = noiseArray.value.filter((tag) => tag !== removedTag);
  };

  const addWhiteNoise = () => {
    if (!sValue.value) {
      Message.error('请输入方差');
      return;
    }

    if (!eValue.value) {
      Message.error('请输入均值');
      return;
    }

    const noiseSignal = `S = ${sValue.value.toString()}, E = ${eValue.value.toString()}`;

    noiseArray.value.push(noiseSignal);

    sValue.value = undefined;
    eValue.value = undefined;
  };

  // const createNoise = () => {};
</script>

<style scoped>
  .function-bottom {
    height: 40%;
    width: 99%;
    margin: 3px 0px;
    display: flex;
    align-items: center;
    align-items: center;
  }
</style>
