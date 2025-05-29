import React from 'react';
import { Container, Row, Col, Card } from 'react-bootstrap';
import { Link } from 'react-router-dom';

const About = () => {

  return (
    <Container className="py-5">
      {/* 页面标题 */}
      <Row className="mb-5">
        <Col className="text-center">
          <h1 className="display-4 fw-bold text-gradient">About Me</h1>
          <div className="mx-auto my-4" style={{
            width: '70px',
            height: '3px',
            background: 'linear-gradient(90deg, #850c62, #f80759)',
            borderRadius: '3px'
          }}></div>
          <p className="lead">Learn more about the person behind this blog</p>
        </Col>
      </Row>

      {/* 个人介绍 */}
      <Row className="mb-5">
        <Col lg={4} className="mb-4 mb-lg-0">
          <Card className="border-0 shadow-sm h-100">
            <Card.Img
              variant="top"
              src="https://via.placeholder.com/400x400?text=Profile+Photo"
              alt="Profile"
              className="rounded-circle mx-auto mt-4"
              style={{ width: '200px', height: '200px', objectFit: 'cover' }}
            />
            <Card.Body className="text-center">
              <Card.Title className="fw-bold fs-4">zhuzhuc</Card.Title>
              <Card.Subtitle className="mb-3 text-muted">Software Developer</Card.Subtitle>
              <div className="d-flex justify-content-center mb-3">
                <a href="https://github.com/zhuzhuc" className="mx-2 text-dark" target="_blank" rel="noopener noreferrer">
                  <i className="fab fa-github fa-lg"></i>
                </a>
              </div>
              <div className="d-grid gap-2">
                <Link to="/contact" className="btn btn-outline-primary">
                  <i className="fas fa-envelope me-2"></i>Contact Me
                </Link>
              </div>
            </Card.Body>
          </Card>
        </Col>
        <Col lg={8}>
          <Card className="border-0 shadow-sm h-100">
            <Card.Body>
              <h2 className="mb-4">My Story</h2>
              <p>
                Hello! I'm a student majoring in software engineering at the University of Science and Technology of China.
              </p>
              {/* <p>
                I specialize in full-stack development, with expertise in Go, React, and cloud technologies.
                This blog is my platform to share knowledge, insights, and experiences from my journey in the tech world.
              </p>
              <p>
                When I'm not coding, you can find me hiking in the mountains, reading science fiction, or experimenting with new recipes in the kitchen.
                I believe in continuous learning and pushing the boundaries of what's possible with technology.
              </p> */}
              <h3 className="mt-4 mb-3">Skills & Expertise</h3>
              <div className="mb-3">
                <h5>Programming Languages</h5>
                <div className="d-flex flex-wrap">
                  {['Go', 'JavaScript', 'TypeScript', 'Python', 'SQL', ].map(skill => (
                    <span key={skill} className="badge bg-primary me-2 mb-2 p-2">{skill}</span>
                  ))}
                </div>
              </div>
              <div className="mb-3">
                <h5>Frameworks & Libraries</h5>
                <div className="d-flex flex-wrap">
                  {['React', 'Fiber', 'Node.js', 'Express', 'Bootstrap', 'Material UI'].map(skill => (
                    <span key={skill} className="badge bg-success me-2 mb-2 p-2">{skill}</span>
                  ))}
                </div>
              </div>
              <div>
                <h5>Tools & Platforms</h5>
                <div className="d-flex flex-wrap">
                  {['Docker', 'Kubernetes', 'AWS', 'Git', 'GitHub', 'VS Code', 'Goland'].map(skill => (
                    <span key={skill} className="badge bg-info text-dark me-2 mb-2 p-2">{skill}</span>
                  ))}
                </div>
              </div>
            </Card.Body>
          </Card>
        </Col>
      </Row>

      {/* 博客目的 */}
      <Row className="mb-5">
        <Col>
          <Card className="border-0 shadow-sm">
            <Card.Body>
              <h2 className="mb-4 text-center">About This Blog</h2>
              <p>
                This blog was created as a platform to share my knowledge, experiences, and insights in the world of software development.
                Here, you'll find articles on Go programming, web development, cloud technologies, and best practices in software engineering.
              </p>
              <p>
                My goal is to create content that is both educational and practical, helping fellow developers solve real-world problems
                and stay up-to-date with the latest trends and technologies in our rapidly evolving industry.
              </p>
              <p>
                I believe in the power of community and knowledge sharing. If you find the content helpful or have suggestions for topics
                you'd like me to cover, please don't hesitate to reach out through the contact page.
              </p>
              <div className="text-center mt-4">
                <Link to="/" className="btn btn-primary me-3">
                  <i className="fas fa-book-open me-2"></i>Read Blog
                </Link>
                <Link to="/contact" className="btn btn-outline-primary">
                  <i className="fas fa-envelope me-2"></i>Get in Touch
                </Link>
              </div>
            </Card.Body>
          </Card>
        </Col>
      </Row>

      {/* 技术栈 */}
      <Row>
        <Col>
          <Card className="border-0 shadow-sm">
            <Card.Body>
              <h2 className="mb-4 text-center">Blog Technology Stack</h2>
              <p className="text-center mb-4">
                This blog is built with modern technologies to provide a fast, responsive, and secure experience.
              </p>
              <Row className="row-cols-1 row-cols-md-3 g-4 text-center">
                <Col>
                  <div className="p-3">
                    <i className="fab fa-react fa-3x mb-3 text-primary"></i>
                    <h4>Frontend</h4>
                    <p>Built with React, Bootstrap, and modern JavaScript to create a responsive and interactive user interface.</p>
                  </div>
                </Col>
                <Col>
                  <div className="p-3">
                    <i className="fas fa-server fa-3x mb-3 text-success"></i>
                    <h4>Backend</h4>
                    <p>Powered by Go and Fiber framework, providing a fast and efficient API for the frontend.</p>
                  </div>
                </Col>
                <Col>
                  <div className="p-3">
                    <i className="fas fa-database fa-3x mb-3 text-info"></i>
                    <h4>Database</h4>
                    <p>Uses MySQL for reliable data storage, with GORM as the ORM for database operations.</p>
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

export default About;