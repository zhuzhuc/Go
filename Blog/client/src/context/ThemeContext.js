import React, { createContext, useState, useEffect, useContext } from 'react';

// 创建主题上下文
const ThemeContext = createContext();

// 创建主题提供者组件
export const ThemeProvider = ({ children }) => {
  // 从 localStorage 获取保存的主题，默认为 'light'
  const [theme, setTheme] = useState(() => {
    const savedTheme = localStorage.getItem('theme');
    return savedTheme || 'light';
  });

  // 从 localStorage 获取保存的粒子效果，默认为 'dream'
  const [particleEffect, setParticleEffect] = useState(() => {
    const savedEffect = localStorage.getItem('particleEffect');
    return savedEffect || 'dream';
  });

  // 切换主题函数
  const toggleTheme = () => {
    setTheme(prevTheme => {
      const newTheme = prevTheme === 'light' ? 'dark' : 'light';
      localStorage.setItem('theme', newTheme);
      return newTheme;
    });
  };

  // 更改粒子效果函数
  const changeParticleEffect = (effect) => {
    console.log(`ThemeContext: 更改粒子效果为 ${effect}`);

    // 检查效果是否有效
    if (!['colorful', 'dream', 'stars'].includes(effect)) {
      console.error(`无效的粒子效果: ${effect}`);
      return;
    }

    // 更新状态
    setParticleEffect(effect);

    // 保存到 localStorage
    try {
      localStorage.setItem('particleEffect', effect);
      console.log(`粒子效果已保存到 localStorage: ${effect}`);
    } catch (error) {
      console.error('保存粒子效果到 localStorage 时出错:', error);
    }
  };

  // 当主题变化时，更新 body 的 data-theme 属性
  useEffect(() => {
    document.body.setAttribute('data-theme', theme);

    // 应用主题相关的类
    if (theme === 'dark') {
      document.body.classList.add('dark-theme');
      document.body.classList.add('dark-mode'); // 添加 dark-mode 类以兼容博客卡片样式
    } else {
      document.body.classList.remove('dark-theme');
      document.body.classList.remove('dark-mode'); // 移除 dark-mode 类
    }
  }, [theme]);

  // 提供上下文值
  const contextValue = {
    theme,
    toggleTheme,
    particleEffect,
    changeParticleEffect
  };

  return (
    <ThemeContext.Provider value={contextValue}>
      {children}
    </ThemeContext.Provider>
  );
};

// 创建自定义 hook 以便于使用 ThemeContext
export const useTheme = () => {
  const context = useContext(ThemeContext);
  if (!context) {
    throw new Error('useTheme 必须在 ThemeProvider 内部使用');
  }
  return context;
};

export default ThemeContext;
