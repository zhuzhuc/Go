import React, { useState } from 'react';
import { Container, Row, Col, Form, Button, Alert, Card } from 'react-bootstrap';
import { Link, useNavigate } from 'react-router-dom';
import axios from 'axios';
import { useAuth } from '../context/AuthContext';
import { useTheme } from '../context/ThemeContext';
import '../styles/Register.css';

const Register = () => {
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    password: '',
    confirmPassword: ''
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const navigate = useNavigate();
  const { login } = useAuth();
  const { theme } = useTheme();

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

    // Validate password match
    if (formData.password !== formData.confirmPassword) {
      setError('Passwords do not match');
      setLoading(false);
      return;
    }

    try {
      // Connect to your backend registration API
      console.log('Sending registration request to:', `${process.env.REACT_APP_API_ROOT}/register`);
      console.log('Registration data:', formData);
      const response = await axios.post(`${process.env.REACT_APP_API_ROOT}/register`, formData);

      if (response.status === 201) {
        // Registration successful, auto login
        login(response.data.user, response.data.token);

        // Show success message
        alert('Registration successful!');

        // Redirect to homepage
        navigate('/');
      }
    } catch (err) {
      console.error('Registration failed:', err);
      setError(err.response?.data?.message || 'Registration failed, please try again later');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Container>
      <Row className="justify-content-center mt-5">
        <Col xs={12} md={8} lg={6}>
          <Card className={`register-card border-0 ${theme === 'dark' ? 'dark-card' : ''}`}>
            <Card.Body className="p-5">
              <div className="text-center mb-4">
                <h2 className="register-title">Create Account</h2>
                <p className="register-subtitle">Join our blog community</p>
              </div>

              {error && (
                <Alert variant="danger" className="mb-4">
                  <i className="fas fa-exclamation-circle me-2"></i>
                  {error}
                </Alert>
              )}

              <Form onSubmit={handleSubmit}>
                <Form.Group className="mb-4">
                  <Form.Label className="form-label">Username</Form.Label>
                  <div className="input-group register-input-group">
                    <span className="input-group-text input-icon-container">
                      <i className="fas fa-user"></i>
                    </span>
                    <Form.Control
                      type="text"
                      name="username"
                      value={formData.username}
                      onChange={handleChange}
                      placeholder="Choose a username"
                      required
                      className="register-form-control shadow-none"
                    />
                  </div>
                </Form.Group>

                <Form.Group className="mb-4">
                  <Form.Label className="form-label">Email</Form.Label>
                  <div className="input-group register-input-group">
                    <span className="input-group-text input-icon-container">
                      <i className="fas fa-envelope"></i>
                    </span>
                    <Form.Control
                      type="email"
                      name="email"
                      value={formData.email}
                      onChange={handleChange}
                      placeholder="Enter your email address"
                      required
                      className="register-form-control shadow-none"
                    />
                  </div>
                </Form.Group>

                <Form.Group className="mb-4">
                  <Form.Label className="form-label">Password</Form.Label>
                  <div className="input-group register-input-group">
                    <span className="input-group-text input-icon-container">
                      <i className="fas fa-lock"></i>
                    </span>
                    <Form.Control
                      type="password"
                      name="password"
                      value={formData.password}
                      onChange={handleChange}
                      placeholder="Create a strong password"
                      required
                      className="register-form-control shadow-none"
                    />
                  </div>
                </Form.Group>

                <Form.Group className="mb-4">
                  <Form.Label className="form-label">Confirm Password</Form.Label>
                  <div className="input-group register-input-group">
                    <span className="input-group-text input-icon-container">
                      <i className="fas fa-lock"></i>
                    </span>
                    <Form.Control
                      type="password"
                      name="confirmPassword"
                      value={formData.confirmPassword}
                      onChange={handleChange}
                      placeholder="Enter your password again"
                      required
                      className="register-form-control shadow-none"
                    />
                  </div>
                </Form.Group>

                <Button
                  type="submit"
                  className="register-button w-100 py-2 mb-4"
                  disabled={loading}
                >
                  {loading ? (
                    <>
                      <span className="spinner-border spinner-border-sm me-2" role="status" aria-hidden="true"></span>
                      Registering...
                    </>
                  ) : (
                    <>Register</>
                  )}
                </Button>

                <div className="text-center">
                  <p className="mb-0 have-account-text">
                    Already have an account? <Link to="/login" className="login-link">Login</Link>
                  </p>
                </div>
              </Form>
            </Card.Body>
          </Card>

          <div className="text-center mt-4">
            <Link to="/" className="return-home">
              <i className="fas fa-arrow-left"></i>
              Return to Home
            </Link>
          </div>
        </Col>
      </Row>
    </Container>
  );
};

export default Register;
