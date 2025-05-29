import React, { useState } from "react";
import { Col, Container, Row } from "react-bootstrap";
import { Link, useNavigate } from "react-router-dom";
import { useAuth } from "../../context/AuthContext";
import { useTheme } from "../../context/ThemeContext";
import "../../styles/Header.css";

const Header = () => {
  const navigate = useNavigate();
  const { isAuthenticated, user, logout } = useAuth();
  const { theme, toggleTheme, particleEffect, changeParticleEffect } = useTheme();

  // Particle effect selector state
  const [showEffectMenu, setShowEffectMenu] = useState(false);

  // Particle effect options
  const effectOptions = [
    { id: 'colorful', name: 'Colorful Particles' },
    { id: 'dream', name: 'Dream Effect' },
    { id: 'stars', name: 'Starry Sky' }
  ];

  return (
    <>
      <Container fluid className="container-fluid header">
        <Row className="align-items-center py-3">
          <Col md={4} className="d-none d-md-block">
            <div className="d-flex align-items-center">
              <Link to="/" className="blog-logo">
                <i className="fas fa-blog fa-2x"></i>
                <span className="fs-4">ZZ_Blog</span>
              </Link>

              {/* Particle effect selector */}
              <div className="ms-4 position-relative">
                <button
                  className="bg-selector-button"
                  onClick={() => setShowEffectMenu(!showEffectMenu)}
                >
                  <i className="fas fa-paint-brush"></i>
                  Background
                  <i className={`fas fa-chevron-${showEffectMenu ? 'up' : 'down'} chevron ${showEffectMenu ? 'open' : ''}`}></i>
                </button>

                {showEffectMenu && (
                  <div className="effect-menu">
                    <div className="p-2">
                      {effectOptions.map((option) => (
                        <button
                          key={option.id}
                          className={`effect-option ${particleEffect === option.id ? 'active' : ''}`}
                          onClick={() => {
                            console.log(`选择粒子效果: ${option.id}`);
                            changeParticleEffect(option.id);
                            console.log(`粒子效果已更改为: ${option.id}`);
                            setShowEffectMenu(false);
                          }}
                        >
                          <i className={`fas fa-${
                            option.id === 'colorful' ? 'palette' :
                            option.id === 'dream' ? 'cloud' : 'star'
                          }`}></i>
                          {option.name}
                          {particleEffect === option.id && (
                            <i className="fas fa-check check"></i>
                          )}
                        </button>
                      ))}
                    </div>
                  </div>
                )}
              </div>
            </div>
          </Col>
          <Col xs={12} md={4} className="d-flex justify-content-center align-items-center">
            {/* 移动设备上显示 Logo */}
            <div className="d-md-none d-flex align-items-center me-3">
              <Link to="/" className="blog-logo">
                <i className="fas fa-blog fa-lg"></i>
                <span>ZZ_Blog</span>
              </Link>
            </div>

            <div className="d-flex flex-wrap justify-content-center">
              <Link to="/" className="nav-button">
                <i className="fas fa-home"></i>
                <span className="d-none d-sm-inline">Home</span>
              </Link>

              <Link to="/about" className="nav-button">
                <i className="fas fa-info-circle"></i>
                <span className="d-none d-sm-inline">About</span>
              </Link>

              <Link to="/contact" className="nav-button">
                <i className="fas fa-envelope"></i>
                <span className="d-none d-sm-inline">Contact</span>
              </Link>

              {isAuthenticated && (
                <Link to="/add" className="nav-button new-post-button">
                  <i className="fas fa-plus-circle"></i>
                  <span className="d-none d-sm-inline">New Post</span>
                </Link>
              )}
            </div>
          </Col>
          <Col md={4} className="d-none d-md-block text-end">
            {isAuthenticated ? (
              <div className="d-flex align-items-center justify-content-end">
                {/* 主题切换按钮 */}
                <div className="me-4">
                  <button
                    className="theme-toggle-btn"
                    onClick={toggleTheme}
                    aria-label={`Switch to ${theme === 'light' ? 'dark' : 'light'} mode`}
                  >
                    <span className="visually-hidden">
                      {theme === 'light' ? 'Switch to dark mode' : 'Switch to light mode'}
                    </span>
                  </button>
                </div>
                <Link to="/profile" className="user-section">
                  <div className="user-avatar">
                    {user?.avatar ? (
                      <img
                        src={user.avatar.startsWith('http') ? user.avatar : `${process.env.REACT_APP_API_ROOT}${user.avatar}`}
                        alt={user.username}
                        className="avatar-image"
                        onError={(e) => {
                          console.log('头像加载失败，使用默认头像');
                          e.target.onerror = null;
                          e.target.src = `https://ui-avatars.com/api/?name=${user.username}&background=random&color=fff&size=128`;
                        }}
                      />
                    ) : (
                      <i className="fas fa-user"></i>
                    )}
                  </div>
                  <span className="username">{user?.username || 'User'}</span>
                </Link>
                <button
                  className="nav-button logout-button"
                  onClick={() => {
                    logout();
                    navigate('/');
                  }}
                >
                  <i className="fas fa-sign-out-alt"></i>
                  <span>Logout</span>
                </button>
              </div>
            ) : (
              <div className="d-flex align-items-center justify-content-end">
                {/* 主题切换按钮 */}
                <div className="me-4">
                  <button
                    className="theme-toggle-btn"
                    onClick={toggleTheme}
                    aria-label={`Switch to ${theme === 'light' ? 'dark' : 'light'} mode`}
                  >
                    <span className="visually-hidden">
                      {theme === 'light' ? 'Switch to dark mode' : 'Switch to light mode'}
                    </span>
                  </button>
                </div>
                <Link to="/login" className="nav-button login-button">
                  <i className="fas fa-sign-in-alt"></i>
                  <span>Login</span>
                </Link>
              </div>
            )}
          </Col>
        </Row>
      </Container>
      {/* 移动设备上的登录/注销按钮 */}
      <div className="d-md-none mobile-auth-section">
        <Container>
          <div className="d-flex justify-content-center align-items-center py-1">
            {isAuthenticated ? (
              <div className="d-flex align-items-center justify-content-center">
                <Link to="/profile" className="d-flex align-items-center">
                  <div className="user-avatar" style={{ width: '28px', height: '28px' }}>
                    {user?.avatar ? (
                      <img
                        src={user.avatar.startsWith('http') ? user.avatar : `${process.env.REACT_APP_API_ROOT}${user.avatar}`}
                        alt={user.username}
                        className="avatar-image"
                        style={{ width: '100%', height: '100%', objectFit: 'cover', borderRadius: '50%' }}
                        onError={(e) => {
                          console.log('头像加载失败，使用默认头像');
                          e.target.onerror = null;
                          e.target.src = `https://ui-avatars.com/api/?name=${user.username}&background=random&color=fff&size=128`;
                        }}
                      />
                    ) : (
                      <i className="fas fa-user fa-sm"></i>
                    )}
                  </div>
                  <span className="username me-3">{user?.username || 'User'}</span>
                </Link>
                <button
                  className="mobile-logout-button"
                  onClick={() => {
                    logout();
                    navigate('/');
                  }}
                >
                  <i className="fas fa-sign-out-alt me-2"></i> Logout
                </button>
              </div>
            ) : (
              <Link to="/login" className="mobile-login-button">
                <i className="fas fa-sign-in-alt me-2"></i> Login
              </Link>
            )}
          </div>
        </Container>
      </div>
    </>
  );
};

export default Header;