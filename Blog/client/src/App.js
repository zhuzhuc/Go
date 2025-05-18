import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import React, { useEffect, useState } from "react";
import "./App.css";
import Home from "./page/Home";
import Blog from "./page/Blog";
import Add from "./page/Add";
import Edit from "./page/Edit";
import Login from "./page/Login";
import Register from "./page/Register";
import Header from "./components/layout/Header";
import Footer from "./components/layout/Footer";
import { AuthProvider } from "./context/AuthContext";

function App() {
  const [particleEffect, setParticleEffect] = useState('dream'); // 默认使用梦幻效果

  // 使用 useMemo 缓存粒子效果配置
  const particleEffects = React.useMemo(() => ({
    colorful: '/particles.json',
    dream: '/particles-dream.json',
    stars: '/particles-stars.json'
  }), []);

  // 初始化粒子效果
  useEffect(() => {
    if (window.particlesJS) {
      // 清除之前的粒子效果实例
      if (window.pJSDom && window.pJSDom.length > 0) {
        window.pJSDom.forEach(dom => dom.pJS.fn.vendors.destroypJS());
        window.pJSDom = [];
      }

      // 加载新的粒子效果
      window.particlesJS.load('particles-js', particleEffects[particleEffect], function() {
        console.log(`粒子效果 ${particleEffect} 已加载`);
      });
    }
  }, [particleEffect, particleEffects]);

  // 粒子效果选择器已移至 Header 组件中

  return (
    <AuthProvider>
      <Router>
        {/* 粒子背景容器 */}
        <div id="particles-js"></div>

        {/* 粒子效果选择器已移至 Header 组件中 */}

        {/* 页面布局 */}
        <div className="app-container">
          <Header particleEffect={particleEffect} setParticleEffect={setParticleEffect} />
          <main className="main-content">
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/blog/:id" element={<Blog />} />
              <Route path="/add" element={<Add />} />
              <Route path="/edit/:id" element={<Edit />} />
              <Route path="/login" element={<Login />} />
              <Route path="/register" element={<Register />} />
            </Routes>
          </main>
          <Footer />
        </div>
      </Router>
    </AuthProvider>
  );
}

export default App;
