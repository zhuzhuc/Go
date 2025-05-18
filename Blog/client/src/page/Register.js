import React, { useState } from 'react';
import { Container, Row, Col, Form, Button, Alert, Card } from 'react-bootstrap';
import { Link, useNavigate } from 'react-router-dom';
import axios from 'axios';
import { useAuth } from '../context/AuthContext';

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
          <Card className="shadow-sm border-0">
            <Card.Body className="p-5">
              <div className="text-center mb-4">
                <h2 className="text-gradient">Create Account</h2>
                <p className="lead text-muted">Join our blog community</p>
              </div>

              {error && (
                <Alert variant="danger" className="mb-4">
                  <i className="fas fa-exclamation-circle me-2"></i>
                  {error}
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
                      placeholder="Choose a username"
                      required
                      className="border-0 shadow-none ps-0"
                    />
                  </div>
                </Form.Group>

                <Form.Group className="mb-4">
                  <Form.Label>Email</Form.Label>
                  <div className="input-group">
                    <span className="input-group-text bg-light border-0">
                      <i className="fas fa-envelope text-muted"></i>
                    </span>
                    <Form.Control
                      type="email"
                      name="email"
                      value={formData.email}
                      onChange={handleChange}
                      placeholder="Enter your email address"
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
                      placeholder="Create a strong password"
                      required
                      className="border-0 shadow-none ps-0"
                    />
                  </div>
                </Form.Group>

                <Form.Group className="mb-4">
                  <Form.Label>Confirm Password</Form.Label>
                  <div className="input-group">
                    <span className="input-group-text bg-light border-0">
                      <i className="fas fa-lock text-muted"></i>
                    </span>
                    <Form.Control
                      type="password"
                      name="confirmPassword"
                      value={formData.confirmPassword}
                      onChange={handleChange}
                      placeholder="Enter your password again"
                      required
                      className="border-0 shadow-none ps-0"
                    />
                  </div>
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
                      Registering...
                    </>
                  ) : (
                    <>Register</>
                  )}
                </Button>

                <div className="text-center">
                  <p className="mb-0">
                    Already have an account? <Link to="/login" className="text-decoration-none">Login</Link>
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

export default Register;
