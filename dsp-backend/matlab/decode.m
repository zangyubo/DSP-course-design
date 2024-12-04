% 用户选择加密后的音频文件
[file, path] = uigetfile('*.wav', '请选择加密后的音频文件');
if isequal(file, 0)
    disp('用户取消了选择');
    return;
end
filename = fullfile(path, file);

% 读取加密后的音频文件
[y_encrypted, Fs] = audioread(filename); % y_encrypted: 加密后的音频信号, Fs: 采样率

% 用户输入密钥
key = input('请输入密钥字符串: ', 's');

% 将密钥转换为 ASCII 值
key_ascii = double(key);

% 解密音频信号 (异或解密)
y_decrypted = y_encrypted;
key_length = length(key_ascii);
for i = 1:length(y_encrypted)
    % 对音频每个采样点进行解密
    y_decrypted(i) = bitxor(y_encrypted(i), key_ascii(mod(i-1, key_length) + 1)); 
end

% 保存解密后的音频文件
decrypted_filename = 'decrypted_audio.wav';
audiowrite(decrypted_filename, y_decrypted, Fs);
disp(['解密音频已保存为: ', decrypted_filename]);

% 播放解密后的音频
disp('播放解密后的音频...');
sound(y_decrypted, Fs);
