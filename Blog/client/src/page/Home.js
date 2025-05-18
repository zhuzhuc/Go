import React, { useState, useEffect } from 'react';
import { Container, Row, Col, Button, Modal } from "react-bootstrap";
import axios from "axios";
import { Link, useNavigate } from 'react-router-dom';
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
import { useAuth } from '../context/AuthContext';

const Home = () => {
  const [apiData, setApiData] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const [showLoginModal, setShowLoginModal] = useState(false);
  const [selectedBlogId, setSelectedBlogId] = useState(null);
  const { isAuthenticated } = useAuth();
  const navigate = useNavigate();

  // Handle read more click
  const handleReadMoreClick = (e, blogId) => {
    if (!isAuthenticated) {
      e.preventDefault();
      setSelectedBlogId(blogId);
      setShowLoginModal(true);
    }
    // If authenticated, the Link will work normally
  };

  // Handle login button click in modal
  const handleLoginClick = () => {
    setShowLoginModal(false);
    navigate('/login');
  };

  useEffect(() => {
    const fetchData = async() => {
      try{
        const apiUrl = process.env.REACT_APP_API_ROOT;
        console.log('请求的 API URL:', apiUrl);
        const response = await axios.get(apiUrl);
        console.log('API 响应数据:', response.data);

        if(response.status === 200){
          if(response?.data.statusText === "OK"){
            setApiData(response?.data?.blog_records);
          } else {
            console.log('响应的 statusText 不是 "OK":', response.data.statusText);
          }
        } else {
          console.log('响应状态码不是 200:', response.status);
        }
      } catch (error) {
        setError(error.message);
        console.error('Error fetching data:', error.message, error.response);
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
    return () => {
      // 可在此添加清理逻辑
    };
  }, []);

  if (isLoading) {
    return (
      <Container>
        <Row>
          <Col xs="12" className="text-center spinner">
            <div className="spinner-border text-primary" role="status">
              <span className="visually-hidden">Loading...</span>
            </div>
          </Col>
        </Row>
      </Container>
    );
  }

  if (error) {
    return (
      <Container>
        <Row>
          <Col xs="12" className="text-center py-5">
            <div className="alert alert-danger" role="alert">
              <i className="fas fa-exclamation-circle me-2"></i>
              Error: {error}
            </div>
          </Col>
        </Row>
      </Container>
    );
  }

  return (
    <Container>
      <Row className="mb-4">
        <Col xs="12" className="py-4">
          <div className="text-center">
            <h1 className="display-4 fw-bold text-gradient">My Blog</h1>
            <p className="lead text-animate">A collection of thoughts, experiences, and ideas</p>

            {/* 添加一个装饰性分隔线 */}
            <div className="mx-auto my-4 text-animate" style={{
              width: '70px',
              height: '3px',
              background: 'linear-gradient(90deg, #850c62, #f80759)',
              borderRadius: '3px',
              animationDelay: '0.3s'
            }}></div>
          </div>
        </Col>
      </Row>

      {/* 博客列表 */}
      <Row className="row-cols-1 row-cols-md-2 row-cols-lg-3 g-4">
        {apiData.length > 0 ? (
          apiData.map((record) => (
            <Col key={record.id}>
              <div className="box h-100">
                <div className="blog-card-content">
                  <div className="title">
                    <Link to={`blog/${record.id}`}>{record.Title}</Link>
                  </div>
                  <div className="author text-muted mb-2">
                    <small><i className="fas fa-user me-1"></i> {record.Author || "Anonymous"}</small>
                  </div>
                  <div className="flex-grow-1 blog-card-excerpt">
                    <ReactMarkdown
                      remarkPlugins={[remarkGfm]}
                      components={{
                        // 不显示图片
                        img: () => null,
                        // 简化链接
                        a: ({children}) => <span>{children}</span>,
                        // 简化段落
                        p: ({children}) => {
                          // 将内容转换为纯文本
                          const textContent = React.Children.toArray(children)
                            .map(child => typeof child === 'string' ? child : '')
                            .join('');

                          // 截断文本
                          const truncatedText = textContent.length > 150
                            ? `${textContent.substring(0, 150)}...`
                            : textContent;

                          return <p>{truncatedText}</p>;
                        }
                      }}
                    >
                      {record.Post}
                    </ReactMarkdown>
                  </div>
                  <div className="mt-auto pt-3 d-flex justify-content-between align-items-center">
                    <div>
                      <Link to={`edit/${record.id}`} className="btn btn-sm btn-outline-secondary me-2">
                        <i className="fas fa-edit"></i>
                      </Link>
                      <button
                        className="btn btn-sm btn-outline-danger"
                        onClick={async (e) => {
                          e.preventDefault();
                          if(window.confirm('Are you sure you want to delete this blog post?')) {
                            try {
                              const apiUrl = process.env.REACT_APP_API_ROOT + "/" + record.id;
                              const response = await axios.delete(apiUrl);
                              if(response.status === 200) {
                                alert('Blog post deleted successfully!');
                                // 刷新博客列表
                                window.location.reload();
                              }
                            } catch (error) {
                              console.error('Error deleting blog post:', error);
                              alert('Failed to delete blog post. Please try again.');
                            }
                          }
                        }}
                      >
                        <i className="fas fa-trash-alt"></i>
                      </button>
                    </div>
                    <Link
                      to={`blog/${record.id}`}
                      className="btn btn-sm btn-outline-primary"
                      onClick={(e) => handleReadMoreClick(e, record.id)}
                    >
                      Read More <i className="fas fa-arrow-right ms-1"></i>
                    </Link>
                  </div>
                </div>
                {record.Image && (
                  <div className="blog-card-image">
                    <img
                      src={process.env.REACT_APP_API_ROOT + record.Image}
                      alt={record.Title}
                      onError={(e) => {
                        e.target.onerror = null;
                        e.target.src = 'https://via.placeholder.com/150x150?text=No+Image';
                      }}
                    />
                  </div>
                )}
              </div>
            </Col>
          ))
        ) : (
          <Col xs="12" className="text-center py-5">
            <div className="alert alert-info" role="alert">
              <i className="fas fa-info-circle me-2"></i>
              No blog posts available yet. Check back soon!
            </div>
          </Col>
        )}
      </Row>

      {/* 添加博客按钮 - 只对已登录用户显示，或者点击时提示登录 */}
      <Row className="mt-5">
        <Col className="text-center">
          {isAuthenticated ? (
            <Link to="/add" className="btn btn-primary">
              <i className="fas fa-plus-circle me-2"></i>
              Add New Blog Post
            </Link>
          ) : (
            <Button
              variant="primary"
              onClick={() => {
                setShowLoginModal(true);
              }}
            >
              <i className="fas fa-plus-circle me-2"></i>
              Add New Blog Post
            </Button>
          )}
        </Col>
      </Row>

      {/* 登录提示模态框 */}
      <Modal show={showLoginModal} onHide={() => setShowLoginModal(false)} centered>
        <Modal.Header closeButton>
          <Modal.Title>登录提示</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <p>您需要登录才能{selectedBlogId ? '查看完整的博客内容' : '发布新博客'}。</p>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={() => setShowLoginModal(false)}>
            取消
          </Button>
          <Button variant="primary" onClick={handleLoginClick}>
            前往登录
          </Button>
        </Modal.Footer>
      </Modal>
    </Container>
  );
};

export default Home;