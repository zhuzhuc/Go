import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import React, { useEffect } from "react";
import "./App.css";
import "./styles/theme.css"; // 导入主题样式
import Home from "./page/Home";
import Blog from "./page/Blog";
import Add from "./page/Add";
import Edit from "./page/Edit";
import Login from "./page/Login";
import Register from "./page/Register";
import Profile from "./page/Profile";
import About from "./page/About";
import Contact from "./page/Contact";
import Debug from "./page/Debug";
import Header from "./components/layout/Header";
import Footer from "./components/layout/Footer";
import { AuthProvider } from "./context/AuthContext";
import { ThemeProvider, useTheme } from "./context/ThemeContext";

function AppContent() {
  // 使用主题上下文
  const { particleEffect, changeParticleEffect } = useTheme();

  // 使用 useMemo 缓存粒子效果配置
  const particleEffects = React.useMemo(() => {
    // 不再需要获取当前网站的基础URL
    // const baseUrl = window.location.origin;

    // 使用相对路径，这样无论部署在哪个域名下都能正常工作
    return {
      colorful: `/particles.json`,
      dream: `/particles-dream.json`,
      stars: `/particles-stars.json`
    };

    // 注释掉使用绝对路径的方式，因为可能导致跨域问题
    // return {
    //   colorful: `${baseUrl}/particles.json`,
    //   dream: `${baseUrl}/particles-dream.json`,
    //   stars: `${baseUrl}/particles-stars.json`
    // };
  }, []);

  // 初始化粒子效果
  useEffect(() => {
    console.log(`尝试加载粒子效果: ${particleEffect}`);

    if (!window.particlesJS) {
      console.error('particlesJS 未加载，请检查 particles.js 库是否正确引入');
      return;
    }

    // 清除之前的粒子效果实例
    if (window.pJSDom && window.pJSDom.length > 0) {
      try {
        window.pJSDom.forEach(dom => dom.pJS.fn.vendors.destroypJS());
        window.pJSDom = [];
        console.log('成功清除之前的粒子效果实例');
      } catch (error) {
        console.error('清除粒子效果实例时出错:', error);
      }
    }

    // 确保粒子效果配置文件路径存在
    const effectPath = particleEffects[particleEffect];
    if (!effectPath) {
      console.error(`未找到粒子效果配置: ${particleEffect}`);
      return;
    }

    console.log(`正在加载粒子效果配置文件: ${effectPath}`);

    // 获取当前效果的配置
    const getParticleConfig = () => {
      if (particleEffect === 'colorful') {
        return {
          particles: {
            number: { value: 100, density: { enable: true, value_area: 1000 } },
            color: { value: ["#850c62", "#f80759", "#6a11cb", "#2575fc", "#00c6ff"] },
            shape: { type: ["circle", "triangle", "star"] },
            opacity: { value: 0.6, random: true },
            size: { value: 4, random: true },
            line_linked: { enable: true, distance: 150, color: "#6a11cb", opacity: 0.3 },
            move: { enable: true, speed: 3, random: true }
          },
          interactivity: {
            detect_on: "window",
            events: {
              onhover: { enable: true, mode: "bubble" },
              onclick: { enable: true, mode: "push" }
            }
          }
        };
      } else if (particleEffect === 'dream') {
        return {
          particles: {
            number: { value: 120, density: { enable: true, value_area: 1200 } },
            color: { value: ["#9C27B0", "#673AB7", "#3F51B5", "#2196F3", "#03A9F4"] },
            shape: { type: "circle" },
            opacity: { value: 0.6, random: true },
            size: { value: 5, random: true },
            line_linked: { enable: true, distance: 150, color: "#9C27B0", opacity: 0.2 },
            move: { enable: true, speed: 1.5, random: true }
          },
          interactivity: {
            detect_on: "window",
            events: {
              onhover: { enable: true, mode: "grab" },
              onclick: { enable: true, mode: "push" }
            }
          }
        };
      } else { // stars
        return {
          particles: {
            number: { value: 160, density: { enable: true, value_area: 1500 } },
            color: { value: "#ffffff" },
            shape: { type: "circle" },
            opacity: { value: 0.7, random: true },
            size: { value: 3, random: true },
            line_linked: { enable: false },
            move: { enable: true, speed: 0.5, random: true }
          },
          interactivity: {
            detect_on: "window",
            events: {
              onhover: { enable: true, mode: "bubble" },
              onclick: { enable: true, mode: "push" }
            }
          }
        };
      }
    };

    // 尝试使用 particlesJS.load 方法加载配置文件
    try {
      // 检查 particlesJS.load 方法是否存在
      if (typeof window.particlesJS.load === 'function') {
        console.log('使用 particlesJS.load 方法加载配置...');
        window.particlesJS.load('particles-js', effectPath, function(response) {
          if (response) {
            console.log(`粒子效果 ${particleEffect} 已成功加载`);
          } else {
            console.error(`粒子效果 ${particleEffect} 加载失败，使用备用方法`);
            // 使用内联配置
            window.particlesJS('particles-js', getParticleConfig());
          }
        });
      } else {
        // 如果 particlesJS.load 方法不存在，直接使用 particlesJS 方法
        console.log('particlesJS.load 方法不可用，使用备用方法');
        window.particlesJS('particles-js', getParticleConfig());
      }
    } catch (error) {
      console.error('加载粒子效果时出错:', error);
      // 出错时也尝试使用备用方法
      try {
        console.log('尝试使用备用方法加载粒子效果...');
        window.particlesJS('particles-js', getParticleConfig());
      } catch (fallbackError) {
        console.error('备用方法也失败:', fallbackError);
      }
    }
  }, [particleEffect, particleEffects]);

  return (
    <Router>
      {/* 粒子背景容器 */}
      <div id="particles-js"></div>

      {/* 页面布局 */}
      <div className="app-container">
        <Header />
        <main className="main-content">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/blog/:id" element={<Blog />} />
            <Route path="/add" element={<Add />} />
            <Route path="/edit/:id" element={<Edit />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/profile" element={<Profile />} />
            <Route path="/about" element={<About />} />
            <Route path="/contact" element={<Contact />} />
            <Route path="/debug" element={<Debug />} />
          </Routes>
        </main>
        <Footer />
      </div>
    </Router>
  );
}

function App() {
  return (
    <AuthProvider>
      <ThemeProvider>
        <AppContent />
      </ThemeProvider>
    </AuthProvider>
  );
}

export default App;
