% 用户选择音频文件
[file, path] = uigetfile({'*.wav;*.mp3', '音频文件 (*.wav, *.mp3)'}, '请选择一个音频文件');
if isequal(file, 0)
    disp('用户取消了选择');
    return;
end
filename = fullfile(path, file);

% 读取音频文件
[y, Fs] = audioread(filename); % y: 音频信号, Fs: 采样率

% 如果音频是多通道，选择一个通道
if size(y, 2) > 1
    y = y(:, 1);
end

% 获取时间轴
t = (0:length(y)-1) / Fs;

% 显示模式选择菜单
mode = input('请输入模式 (1: 单频噪声, 2: 多频噪声, 3: 白噪声, 4: 混响处理): ');

% 根据模式处理音频信号
switch mode
    case 1 % 单频噪声（正弦干扰）
        f = 500; % 噪声频率，单位Hz
        noise = 0.1 * sin(2 * pi * f * t)'; % 生成正弦噪声
        y_processed = y + noise; % 添加噪声

    case 2 % 多频噪声（多正弦干扰）
        f1 = 500; f2 = 1000; % 两个噪声频率
        noise = 0.1 * (sin(2 * pi * f1 * t)' + sin(2 * pi * f2 * t)'); % 生成多频噪声
        y_processed = y + noise; % 添加噪声

    case 3 % 白噪声
        noise = 0.1 * randn(size(y)); % 生成白噪声
        y_processed = y + noise; % 添加噪声

    case 4 % 混响处理
        reverb = reverberator('PreDelay', 50, 'WetDryMix', 0.3); % 创建混响效果
        y_processed = reverb(y); % 应用混响处理

    otherwise
        disp('无效的模式');
        return;
end

% 绘制原始和处理后信号
figure;
subplot(2, 1, 1);
plot(t, y);
title('原始音频信号');
xlabel('时间 (s)');
ylabel('振幅');
grid on;

subplot(2, 1, 2);
plot(t, y_processed);
title('处理后音频信号');
xlabel('时间 (s)');
ylabel('振幅');
grid on;

% 播放原始和处理后音频
disp('播放原始音频...');
sound(y, Fs);
pause(length(y) / Fs + 1); % 等待播放完成

disp('播放处理后音频...');
sound(y_processed, Fs);
