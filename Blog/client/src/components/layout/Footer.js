import React from "react";
import { Col, Container, Row } from "react-bootstrap";
import { Link } from "react-router-dom";

const Footer = () => {
  const currentYear = new Date().getFullYear();

  return (
    <Container fluid className="container-fluid footer">
      <Row className="py-4">
        <Col md="4" className="mb-4 mb-md-0">
          <h3>Navigation</h3>
          <ul className="list-unstyled">
            <li>
              <Link to="/">
                <i className="fas fa-home me-2"></i>Home
              </Link>
            </li>
            <li>
              <Link to="/">
                <i className="fas fa-blog me-2"></i>Blog
              </Link>
            </li>
            <li>
              <Link to="/about">
                <i className="fas fa-info-circle me-2"></i>About
              </Link>
            </li>
            <li>
              <Link to="/contact">
                <i className="fas fa-envelope me-2"></i>Contact
              </Link>
            </li>
          </ul>
        </Col>

        <Col md="4" className="mb-4 mb-md-0">
          <h3>Recent Posts</h3>
          <ul className="list-unstyled">
            <li>
              <Link to="/">
                <i className="fas fa-file-alt me-2"></i>Getting Started with Go
              </Link>
            </li>
            <li>
              <Link to="/">
                <i className="fas fa-file-alt me-2"></i>Building APIs with Fiber
              </Link>
            </li>
            <li>
              <Link to="/">
                <i className="fas fa-file-alt me-2"></i>React and Go Integration
              </Link>
            </li>
          </ul>
        </Col>

        <Col md="4">
          <h3>Connect With Me</h3>
          <div className="social-links pt-2">
            <a
              target="_blank"
              rel="noopener noreferrer"
              href="https://github.com/zhuzhuc"
              className="me-3"
              title="Connect on Github"
            >
              <i className="fab fa-github fa-2x"></i>
            </a>
            <a
              target="_blank"
              rel="noopener noreferrer"
              href="https://twitter.com/"
              className="me-3"
              title="Connect on Twitter"
            >
              <i className="fab fa-twitter fa-2x"></i>
            </a>
            <a
              target="_blank"
              rel="noopener noreferrer"
              href="https://linkedin.com/"
              className="me-3"
              title="Connect on LinkedIn"
            >
              <i className="fab fa-linkedin fa-2x"></i>
            </a>
          </div>

          <div className="mt-4">
            <h5>Subscribe to Newsletter</h5>
            <div className="input-group">
              <input type="email" className="form-control" placeholder="Your email" />
              <button className="btn btn-light" type="button">
                <i className="fas fa-paper-plane"></i>
              </button>
            </div>
          </div>
        </Col>
      </Row>

      <Row className="border-top pt-3 mt-3">
        <Col className="text-center">
          <p className="mb-0">
            &copy; {currentYear} Personal Blog. All rights reserved.
          </p>
        </Col>
      </Row>
    </Container>
  );
};

export default Footer;