import React, { useEffect } from "react";
import { Col, Container, Row } from "react-bootstrap";
import { Link } from "react-router-dom";
import axios from "axios";
import { useTheme } from "../../context/ThemeContext";
import "./Footer.css";

const Footer = () => {
  const { theme } = useTheme();
  // 获取最新博客文章并更新 Recent Posts 部分
  useEffect(() => {
    const fetchRecentPosts = async () => {
      try {
        const apiUrl = process.env.REACT_APP_API_ROOT;
        const response = await axios.get(apiUrl);

        if (response.status === 200 && response.data.statusText === "OK") {
          const posts = response.data.blog_records || [];
          // 按创建时间排序，获取最新的3篇文章
          const recentPosts = [...posts]
            .sort((a, b) => new Date(b.CreatedAt) - new Date(a.CreatedAt))
            .slice(0, 3);

          // 获取 Recent Posts 占位符元素
          const placeholder = document.getElementById('recent-posts-placeholder');
          if (placeholder && recentPosts.length > 0) {
            // 创建文章列表
            const ul = document.createElement('ul');
            ul.className = 'list-unstyled';

            recentPosts.forEach(post => {
              const li = document.createElement('li');
              li.className = 'mb-3';

              const link = document.createElement('a');
              link.href = `/blog/${post.id}`;

              // 截断标题，如果太长
              const title = post.Title.length > 30
                ? post.Title.substring(0, 30) + '...'
                : post.Title;

              link.innerHTML = `<i class="fas fa-file-alt me-2" style="color: var(--accent-color)"></i>${title}`;

              li.appendChild(link);
              ul.appendChild(li);
            });

            // 替换占位符内容
            placeholder.innerHTML = '';
            placeholder.appendChild(ul);
          } else if (placeholder && recentPosts.length === 0) {
            placeholder.innerHTML = '<p class="text-secondary"><i class="fas fa-info-circle me-2" style="color: var(--accent-color)"></i>No posts available yet.</p>';
          }
        }
      } catch (error) {
        console.error('Error fetching recent posts:', error);
        const placeholder = document.getElementById('recent-posts-placeholder');
        if (placeholder) {
          placeholder.innerHTML = '<p class="text-secondary"><i class="fas fa-exclamation-circle me-2" style="color: var(--accent-color)"></i>Failed to load recent posts.</p>';
        }
      }
    };

    fetchRecentPosts();
  }, []);

  return (
    <footer className={`footer mt-auto ${theme === 'dark' ? 'dark-footer' : ''}`}>
      <div className="footer-main py-1">
        <Container>
          <Row className="py-1">
            <Col lg="3" md="5" className="mb-4 mb-lg-0 fade-in">
              <h4 className="footer-heading mb-3">Quick Links</h4>
              <ul className="list-unstyled footer-links">
                <li className="mb-2">
                  <Link to="/" className="d-flex align-items-center">
                    <i className="fas fa-home me-1 small"></i>
                    <span>Home</span>
                  </Link>
                </li>
                <li className="mb-2">
                  <Link to="/about" className="d-flex align-items-center">
                    <i className="fas fa-info-circle me-2 small"></i>
                    <span>About</span>
                  </Link>
                </li>
                <li className="mb-2">
                  <Link to="/contact" className="d-flex align-items-center">
                    <i className="fas fa-envelope me-2 small"></i>
                    <span>Contact</span>
                  </Link>
                </li>
                <li className="mb-2">
                  <Link to="/login" className="d-flex align-items-center">
                    <i className="fas fa-sign-in-alt me-2 small"></i>
                    <span>Login</span>
                  </Link>
                </li>
              </ul>
            </Col>

            <Col lg="5" md="6" className="mb-4 mb-lg-0 fade-in" style={{ animationDelay: '0.1s' }}>
              <h4 className="footer-heading mb-3">Recent Posts</h4>
              <div id="recent-posts-placeholder" className="py-2">
                <div className="d-flex align-items-center">
                  <div className="spinner-border spinner-border-sm me-2 text-accent" role="status" style={{
                    width: '0.75rem',
                    height: '0.75rem',
                    color: 'var(--accent-color)'
                  }}>
                    <span className="visually-hidden">Loading...</span>
                  </div>
                  <span>Loading recent posts...</span>
                </div>
              </div>
            </Col>

            <Col lg="4" md="12" className="fade-in" style={{ animationDelay: '0.2s' }}>
              <h4 className="footer-heading mb-3">Stay Connected</h4>

              <div className="social-links d-flex mb-3">
                <a
                  target="_blank"
                  rel="noopener noreferrer"
                  href="https://github.com/zhuzhuc"
                  className="social-icon"
                  title="Connect on Github"
                >
                  <i className="fab fa-github"></i>
                </a>
              </div>

              <div className="newsletter">
                <h5 className="mb-3 newsletter-title">Subscribe to Newsletter</h5>
                <div className="input-group">
                  <input
                    type="email"
                    className="form-control newsletter-input"
                    placeholder="Your email"
                  />
                  <button
                    className="btn subscribe-btn"
                    type="button"
                  >
                    <i className="fas fa-paper-plane me-1" style={{ color: 'white' }}></i>
                    Subscribe
                  </button>
                </div>
                <p className="mt-2 small privacy-text">
                  <i className="fas fa-lock me-1"></i>
                  We respect your privacy and will never share your email.
                </p>
              </div>
            </Col>
          </Row>
        </Container>
      </div>

      <div className="footer-bottom py-1">
        <Container>
          <div className="d-flex flex-column flex-md-row justify-content-between align-items-center">
            <div className="d-flex align-items-center mb-3 mb-md-0">
              <div className="me-2 copyright-icon">
                <i className="fas fa-code"></i>
              </div>
              <p className="mb-0 small copyright-text">
                &copy; {new Date().getFullYear()} ZZ_Blog. All rights reserved.
              </p>
            </div>
            <div className="d-flex align-items-center">
              <Link to="/privacy" className="me-4 small d-flex align-items-center hover-effect footer-link">
                <i className="fas fa-shield-alt me-1 small"></i>
                <span>Privacy</span>
              </Link>
              <Link to="/terms" className="small d-flex align-items-center hover-effect footer-link">
                <i className="fas fa-file-contract me-1 small"></i>
                <span>Terms</span>
              </Link>
            </div>
          </div>
        </Container>
      </div>
    </footer>
  );
};

export default Footer;