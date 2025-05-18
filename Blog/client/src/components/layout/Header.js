import React, { useState } from "react";
import { Col, Container, Row } from "react-bootstrap";
import { Link, useNavigate } from "react-router-dom";
import { useAuth } from "../../context/AuthContext";

const Header = ({ particleEffect, setParticleEffect }) => {
  const navigate = useNavigate();
  const { isAuthenticated, user, logout } = useAuth();

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
        <Row className="align-items-center">
          <Col md={4} className="d-none d-md-block">
            <div className="d-flex align-items-center">
              <Link to="/" className="text-white text-decoration-none">
                <i className="fas fa-blog fa-2x me-2"></i>
                <span className="fw-bold">ZZ_Blog</span>
              </Link>

              {/* Particle effect selector */}
              <div className="ms-3 position-relative">
                <button
                  className="btn btn-sm text-white"
                  style={{
                    background: 'rgba(255,255,255,0.2)',
                    borderRadius: '4px',
                    padding: '4px 10px'
                  }}
                  onClick={() => setShowEffectMenu(!showEffectMenu)}
                >
                  <i className="fas fa-paint-brush me-1"></i>
                  Background
                  <i className={`fas fa-chevron-${showEffectMenu ? 'up' : 'down'} ms-1`}></i>
                </button>

                {showEffectMenu && (
                  <div
                    className="position-absolute mt-1"
                    style={{
                      background: 'white',
                      borderRadius: '4px',
                      boxShadow: '0 2px 10px rgba(0,0,0,0.2)',
                      zIndex: 1000,
                      width: '130px',
                      animation: 'fadeIn 0.2s ease-out'
                    }}
                  >
                    {effectOptions.map(option => (
                      <button
                        key={option.id}
                        className="d-block w-100 text-start border-0 bg-transparent py-2 px-3"
                        style={{
                          cursor: 'pointer',
                          borderBottom: option.id !== 'stars' ? '1px solid #eee' : 'none',
                          color: particleEffect === option.id ? '#e91e63' : '#333',
                          fontWeight: particleEffect === option.id ? 'bold' : 'normal'
                        }}
                        onClick={() => {
                          setParticleEffect(option.id);
                          setShowEffectMenu(false);
                        }}
                      >
                        {option.name}
                      </button>
                    ))}
                  </div>
                )}
              </div>
            </div>
          </Col>
          <Col xs={12} md={4}>
            <h1 className="text-center text-uppercase m-0">
              Personal Blog
            </h1>
          </Col>
          <Col md={4} className="d-none d-md-block text-end">
            {isAuthenticated ? (
              <div className="text-white d-flex align-items-center justify-content-end">
                <div className="me-3">
                  <i className="fas fa-user-circle me-2"></i>
                  {user?.username || 'User'}
                </div>
                <button
                  className="btn btn-outline-light btn-sm"
                  onClick={() => {
                    logout();
                    navigate('/');
                  }}
                >
                  <i className="fas fa-sign-out-alt me-1"></i>
                  Logout
                </button>
              </div>
            ) : (
              <Link to="/login" className="btn btn-outline-light btn-sm">
                <i className="fas fa-sign-in-alt me-2"></i>
                Login
              </Link>
            )}
          </Col>
        </Row>
      </Container>
      <Container>
        <nav className="py-2">
          <ul className="menu text-center">
            <li>
              <Link to="/">
                <i className="fas fa-home me-1"></i> Home
              </Link>
            </li>
            <li>
              <Link to="/">
                <i className="fas fa-blog me-1"></i> Blog
              </Link>
            </li>
            <li>
              <Link to="/about">
                <i className="fas fa-info-circle me-1"></i> About
              </Link>
            </li>
            <li>
              <Link to="/contact">
                <i className="fas fa-envelope me-1"></i> Contact
              </Link>
            </li>
            <li className="d-md-none">
              {isAuthenticated ? (
                <>
                  <i className="fas fa-user-circle me-1"></i>
                  {user?.username || 'User'}
                  <button
                    className="btn btn-link text-decoration-none p-0 ms-2"
                    onClick={() => {
                      logout();
                      navigate('/');
                    }}
                  >
                    <i className="fas fa-sign-out-alt me-1"></i> Logout
                  </button>
                </>
              ) : (
                <Link to="/login">
                  <i className="fas fa-sign-in-alt me-1"></i> Login
                </Link>
              )}
            </li>
          </ul>
        </nav>
      </Container>
    </>
  );
};

export default Header;