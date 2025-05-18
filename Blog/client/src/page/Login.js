import React, { useState, useEffect } from 'react';
import { Container, Row, Col, Form, Button, Alert, Card } from 'react-bootstrap';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import axios from 'axios';
import { useAuth } from '../context/AuthContext';

const Login = () => {
  const [formData, setFormData] = useState({
    username: '',
    password: ''
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [redirectMessage, setRedirectMessage] = useState('');
  const navigate = useNavigate();
  const location = useLocation();
  const { login } = useAuth();

  // 检查是否有重定向信息
  useEffect(() => {
    const { state } = location;
    if (state && state.from) {
      setRedirectMessage(state.message || `您需要登录才能访问 ${state.from}`);
    }
  }, [location]);

  // Handle form input changes
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prevState => ({
      ...prevState,
      [name]: value
    }));
  };

  // Handle form submission
  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    try {
      // Connect to your backend login API
      // This is example code, adjust according to your actual API
      const response = await axios.post(`${process.env.REACT_APP_API_ROOT}/login`, formData);

      if (response.status === 200) {
        console.log("Login response:", response.data);
        console.log("Token from response:", response.data.token);

        // 确保 token 存在
        if (!response.data.token) {
          setError('服务器返回的令牌无效');
          setLoading(false);
          return;
        }

        // 手动保存 token 到 localStorage
        localStorage.setItem('token', response.data.token);
        localStorage.setItem('user', JSON.stringify(response.data.user));

        // 使用 AuthContext 的 login 函数
        login(response.data.user, response.data.token);

        // 验证 token 是否正确保存
        console.log("Token saved to localStorage:", localStorage.getItem('token'));

        // Show success message
        alert('Login successful!');

        // 如果有重定向路径，则重定向到该路径，否则重定向到首页
        const { state } = location;
        if (state && state.from) {
          navigate(state.from);
        } else {
          navigate('/');
        }
      }
    } catch (err) {
      console.error('Login failed:', err);
      setError(err.response?.data?.message || 'Login failed, please check your credentials');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Container>
      <Row className="justify-content-center mt-5">
        <Col xs={12} md={8} lg={6}>
          <Card className="shadow-sm border-0">
            <Card.Body className="p-5">
              <div className="text-center mb-4">
                <h2 className="text-gradient">Welcome Back</h2>
                <p className="lead text-muted">Login to your account</p>
              </div>

              {error && (
                <Alert variant="danger" className="mb-4">
                  <i className="fas fa-exclamation-circle me-2"></i>
                  {error}
                </Alert>
              )}

              {redirectMessage && (
                <Alert variant="info" className="mb-4">
                  <i className="fas fa-info-circle me-2"></i>
                  {redirectMessage}
                </Alert>
              )}

              <Form onSubmit={handleSubmit}>
                <Form.Group className="mb-4">
                  <Form.Label>Username</Form.Label>
                  <div className="input-group">
                    <span className="input-group-text bg-light border-0">
                      <i className="fas fa-user text-muted"></i>
                    </span>
                    <Form.Control
                      type="text"
                      name="username"
                      value={formData.username}
                      onChange={handleChange}
                      placeholder="Enter your username"
                      required
                      className="border-0 shadow-none ps-0"
                    />
                  </div>
                </Form.Group>

                <Form.Group className="mb-4">
                  <Form.Label>Password</Form.Label>
                  <div className="input-group">
                    <span className="input-group-text bg-light border-0">
                      <i className="fas fa-lock text-muted"></i>
                    </span>
                    <Form.Control
                      type="password"
                      name="password"
                      value={formData.password}
                      onChange={handleChange}
                      placeholder="Enter your password"
                      required
                      className="border-0 shadow-none ps-0"
                    />
                  </div>
                </Form.Group>

                <Form.Group className="mb-4 d-flex justify-content-between align-items-center">
                  <Form.Check
                    type="checkbox"
                    label="Remember me"
                    id="remember-me"
                  />
                  <Link to="/forgot-password" className="text-decoration-none text-primary">
                    Forget your password?
                  </Link>
                </Form.Group>

                <Button
                  variant="primary"
                  type="submit"
                  className="w-100 py-2 mb-4"
                  disabled={loading}
                >
                  {loading ? (
                    <>
                      <span className="spinner-border spinner-border-sm me-2" role="status" aria-hidden="true"></span>
                      Logging in...
                    </>
                  ) : (
                    <>Login</>
                  )}
                </Button>

                <div className="text-center">
                  <p className="mb-0">
                    No Account? <Link to="/register" className="text-decoration-none">Register</Link>
                  </p>
                </div>
              </Form>
            </Card.Body>
          </Card>

          <div className="text-center mt-4">
            <Link to="/" className="text-decoration-none">
              <i className="fas fa-arrow-left me-2"></i>
              Return to Home
            </Link>
          </div>
        </Col>
      </Row>
    </Container>
  );
};

export default Login;