import React, { useState } from 'react';
import { Container, Row, Col, Form, Button, Card, Alert } from 'react-bootstrap';
import { Link } from 'react-router-dom';

const Contact = () => {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    subject: '',
    message: ''
  });
  const [validated, setValidated] = useState(false);
  const [submitStatus, setSubmitStatus] = useState(null);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prevState => ({
      ...prevState,
      [name]: value
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    const form = e.currentTarget;

    if (form.checkValidity() === false) {
      e.stopPropagation();
      setValidated(true);
      return;
    }

    setValidated(true);

    // Simulate form submission
    setTimeout(() => {
      setSubmitStatus({
        type: 'success',
        message: 'Your message has been sent successfully! I will get back to you soon.'
      });

      // Reset form
      setFormData({
        name: '',
        email: '',
        subject: '',
        message: ''
      });
      setValidated(false);

      // Clear success message after 5 seconds
      setTimeout(() => {
        setSubmitStatus(null);
      }, 5000);
    }, 1000);
  };

  return (
    <Container className="py-5">
      {/* 页面标题 */}
      <Row className="mb-5">
        <Col className="text-center">
          <h1 className="display-4 fw-bold text-gradient">Contact Me</h1>
          <div className="mx-auto my-4" style={{
            width: '70px',
            height: '3px',
            background: 'linear-gradient(90deg, #850c62, #f80759)',
            borderRadius: '3px'
          }}></div>
          <p className="lead">Have a question or want to work together? Reach out to me!</p>
        </Col>
      </Row>

      {/* 联系表单和信息 */}
      <Row>
        <Col lg={7} className="mb-4 mb-lg-0">
          <Card className="border-0 shadow-sm h-100">
            <Card.Body className="p-4">
              <h2 className="mb-4">Send a Message</h2>

              {submitStatus && (
                <Alert variant={submitStatus.type} className="mb-4">
                  {submitStatus.type === 'success' && <i className="fas fa-check-circle me-2"></i>}
                  {submitStatus.type === 'danger' && <i className="fas fa-exclamation-circle me-2"></i>}
                  {submitStatus.message}
                </Alert>
              )}

              <Form noValidate validated={validated} onSubmit={handleSubmit}>
                <Row>
                  <Col md={6} className="mb-3">
                    <Form.Group controlId="contactName">
                      <Form.Label>Your Name</Form.Label>
                      <Form.Control
                        type="text"
                        name="name"
                        value={formData.name}
                        onChange={handleChange}
                        placeholder="Enter your name"
                        required
                      />
                      <Form.Control.Feedback type="invalid">
                        Please provide your name.
                      </Form.Control.Feedback>
                    </Form.Group>
                  </Col>
                  <Col md={6} className="mb-3">
                    <Form.Group controlId="contactEmail">
                      <Form.Label>Email Address</Form.Label>
                      <Form.Control
                        type="email"
                        name="email"
                        value={formData.email}
                        onChange={handleChange}
                        placeholder="Enter your email"
                        required
                      />
                      <Form.Control.Feedback type="invalid">
                        Please provide a valid email address.
                      </Form.Control.Feedback>
                    </Form.Group>
                  </Col>
                </Row>
                <Form.Group className="mb-3" controlId="contactSubject">
                  <Form.Label>Subject</Form.Label>
                  <Form.Control
                    type="text"
                    name="subject"
                    value={formData.subject}
                    onChange={handleChange}
                    placeholder="What is this regarding?"
                    required
                  />
                  <Form.Control.Feedback type="invalid">
                    Please provide a subject.
                  </Form.Control.Feedback>
                </Form.Group>
                <Form.Group className="mb-4" controlId="contactMessage">
                  <Form.Label>Message</Form.Label>
                  <Form.Control
                    as="textarea"
                    name="message"
                    value={formData.message}
                    onChange={handleChange}
                    rows={6}
                    placeholder="Your message here..."
                    required
                    style={{ resize: 'none' }}
                  />
                  <Form.Control.Feedback type="invalid">
                    Please provide your message.
                  </Form.Control.Feedback>
                </Form.Group>
                <div className="d-grid">
                  <Button variant="primary" type="submit" size="lg">
                    <i className="fas fa-paper-plane me-2"></i>Send Message
                  </Button>
                </div>
              </Form>
            </Card.Body>
          </Card>
        </Col>
        <Col lg={5}>
          <Card className="border-0 shadow-sm h-100">
            <Card.Body className="p-4">
              <h2 className="mb-4">Contact Information</h2>
              <p className="mb-4">
                Feel free to reach out to me with any questions, project inquiries, or just to say hello.
                I'm always open to discussing new projects, creative ideas, or opportunities to be part of your vision.
              </p>

              <div className="mb-4">
                <h5 className="mb-3">Connect With Me</h5>
                <div className="d-flex mb-3">
                  <div className="me-3">
                    <i className="fas fa-envelope fa-fw text-primary"></i>
                  </div>
                  <div>
                    <h6 className="mb-1">Email</h6>
                    <p className="mb-0">zzczhuzhu@gmail.com</p>
                  </div>
                </div>
                <div className="d-flex mb-3">
                  <div className="me-3">
                    <i className="fas fa-map-marker-alt fa-fw text-primary"></i>
                  </div>
                  <div>
                    <h6 className="mb-1">Location</h6>
                    <p className="mb-0">China</p>
                  </div>
                </div>
                <div className="d-flex">
                  <div className="me-3">
                    <i className="fas fa-clock fa-fw text-primary"></i>
                  </div>
                  <div>
                    <h6 className="mb-1">Working Hours</h6>
                    <p className="mb-0">9AM - 5PM PST</p>
                  </div>
                </div>
              </div>

              <div className="mb-4">
                <h5 className="mb-3">Follow Me</h5>
                <div className="d-flex">
                  <a href="https://github.com/zhuzhuc" className="me-3 text-dark" target="_blank" rel="noopener noreferrer">
                    <i className="fab fa-github fa-2x"></i>
                  </a>
            
                </div>
              </div>

              <div>
                <h5 className="mb-3">Looking for my blog?</h5>
                <Link to="/" className="btn btn-outline-primary">
                  <i className="fas fa-book-open me-2"></i>Visit Blog
                </Link>
              </div>
            </Card.Body>
          </Card>
        </Col>
      </Row>

      {/* FAQ 部分 */}
      <Row className="mt-5">
        <Col>
          <Card className="border-0 shadow-sm">
            <Card.Body className="p-4">
              <h2 className="text-center mb-4">Frequently Asked Questions</h2>
              <Row className="row-cols-1 row-cols-md-2 g-4">
                <Col>
                  <div className="mb-4">
                    <h5><i className="fas fa-question-circle text-primary me-2"></i>What services do you offer?</h5>
                    <p className="mb-0">
                      I specialize in full-stack development, focusing on Go for backend and React for frontend.
                      I can help with web applications, APIs, database design, and cloud deployment.
                    </p>
                  </div>
                </Col>
                <Col>
                  <div className="mb-4">
                    <h5><i className="fas fa-question-circle text-primary me-2"></i>Do you take on freelance projects?</h5>
                    <p className="mb-0">
                      Yes, I'm open to freelance opportunities. Feel free to contact me with details about your project,
                      and we can discuss how I can help bring your vision to life.
                    </p>
                  </div>
                </Col>
                <Col>
                  <div className="mb-4">
                    <h5><i className="fas fa-question-circle text-primary me-2"></i>How quickly do you respond to inquiries?</h5>
                    <p className="mb-0">
                      I typically respond to all inquiries within 24-48 hours during business days.
                      For urgent matters, please indicate so in your message.
                    </p>
                  </div>
                </Col>
                <Col>
                  <div className="mb-4">
                    <h5><i className="fas fa-question-circle text-primary me-2"></i>Can I guest post on your blog?</h5>
                    <p className="mb-0">
                      I'm open to high-quality guest posts that align with the blog's focus on software development.
                      Please contact me with your proposed topic and outline.
                    </p>
                  </div>
                </Col>
              </Row>
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default Contact;