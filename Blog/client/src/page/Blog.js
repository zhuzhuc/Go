import axios from "axios";
import { useParams, Link, useNavigate } from 'react-router-dom';
import React, {useEffect, useState } from 'react';
import { Col, Container, Row } from "react-bootstrap";
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
import { useAuth } from '../context/AuthContext';

// Define the Blog component
const Blog = () => {
    const params = useParams();
    const navigate = useNavigate();
    const [apiData, setApiData] = useState(null);
    const [error, setError] = useState(null);
    const [isLoading, setIsLoading] = useState(true);
    const { isAuthenticated, user } = useAuth();

    // 检查当前用户是否是博客作者
    const isAuthor = () => {
        if (!isAuthenticated || !user || !apiData) return false;
        return apiData.AuthorID === user.id;
    };

    useEffect(() => {
        const fetchData = async() => {
            try{
                const apiUrl = process.env.REACT_APP_API_ROOT + "/" + params.id;
                console.log('请求的 API URL:', apiUrl);
                const response = await axios.get(apiUrl);
                console.log('API 响应数据:', response.data);

                if(response.status === 200){
                    if(response?.data.statusText === "OK"){
                        setApiData(response?.data?.record);
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

        };
    }, [params.id]);

    console.log(apiData)

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

    if (!apiData) {
        return (
            <Container>
                <Row>
                    <Col xs="12" className="text-center py-5">
                        <div className="alert alert-warning" role="alert">
                            <i className="fas fa-info-circle me-2"></i>
                            Cannot find the blog post you're looking for.
                        </div>
                    </Col>
                </Row>
            </Container>
        );
    }

    return (
        <Container>
            <Row className="justify-content-center">
                <Col xs="12" md="10" lg="8">
                    <div className="blog-detail fade-in">
                        {/* 显示特色图片 */}
                        {apiData.Image && (
                            <div className="featured-image mb-5 zoom-in">
                                <img
                                    src={process.env.REACT_APP_API_ROOT + apiData.Image}
                                    alt={apiData.Title}
                                    className="img-fluid rounded"
                                    style={{ width: '100%', height: '400px', objectFit: 'cover' }}
                                    onError={(e) => {
                                        console.log("Image load error:", e);
                                        e.target.onerror = null;
                                        e.target.src = 'https://via.placeholder.com/1200x400?text=Image+Not+Found';
                                    }}
                                />
                            </div>
                        )}

                        <h1 className="slide-in-left">{apiData.Title}</h1>

                        <div className="meta slide-in-left" style={{ animationDelay: '0.1s' }}>
                            <div>
                                <i className="far fa-calendar-alt"></i>
                                {new Date().toLocaleDateString()}
                            </div>
                            <div>
                                <i className="fas fa-user"></i>
                                {apiData.Author || "Anonymous"}
                            </div>
                        </div>

                        {/* 使用 Markdown 渲染博客内容 */}
                        <div className="blog-content slide-in-left" style={{ animationDelay: '0.2s' }}>
                            <ReactMarkdown
                                remarkPlugins={[remarkGfm]}
                                components={{
                                    // 自定义图片渲染，处理相对路径
                                    img: ({node, ...props}) => {
                                        const src = props.src;
                                        // 如果是相对路径，添加API根路径
                                        if (src && !src.startsWith('http') && !src.startsWith('/')) {
                                            props.src = process.env.REACT_APP_API_ROOT + '/static/uploads/' + src;
                                        } else if (src && src.startsWith('/')) {
                                            props.src = process.env.REACT_APP_API_ROOT + src;
                                        }
                                        return (
                                            <div className="text-center my-4">
                                                <img
                                                    {...props}
                                                    className="img-fluid rounded shadow-sm"
                                                    alt={props.alt || ''}
                                                    style={{ maxWidth: '100%', height: 'auto' }}
                                                    onError={(e) => {
                                                        console.log("Image load error:", e);
                                                        e.target.onerror = null;
                                                        e.target.src = 'https://via.placeholder.com/800x400?text=Image+Not+Found';
                                                    }}
                                                />
                                            </div>
                                        );
                                    },
                                    // 自定义链接渲染
                                    a: ({node, children, ...props}) => {
                                        return <a {...props} target="_blank" rel="noopener noreferrer" className="text-decoration-none">{children}</a>
                                    },
                                    // 自定义段落渲染
                                    p: ({node, ...props}) => {
                                        return <p {...props} className="mb-3" style={{ fontSize: '1.1rem', lineHeight: '1.8' }} />
                                    }
                                }}
                            >
                                {apiData.Post}
                            </ReactMarkdown>
                        </div>

                        <div className="mt-5 pt-4 border-top d-flex justify-content-between align-items-center slide-in-left" style={{ animationDelay: '0.3s' }}>
                            <Link to="/" className="btn btn-outline-primary">
                                <i className="fas fa-arrow-left me-2"></i>
                                Back to Blog List
                            </Link>
                            {isAuthor() && (
                                <div className="d-flex">
                                    <Link to={`/edit/${apiData.id}`} className="btn btn-outline-secondary me-3">
                                        <i className="fas fa-edit me-1"></i>
                                        Edit
                                    </Link>
                                    <button
                                        className="btn btn-outline-danger"
                                        onClick={async () => {
                                            if(window.confirm('Are you sure you want to delete this blog post?')) {
                                                try {
                                                    const apiUrl = process.env.REACT_APP_API_ROOT + "/" + params.id;
                                                    const response = await axios.delete(apiUrl, {
                                                        headers: {
                                                            'Authorization': `Bearer ${localStorage.getItem('token')}`
                                                        }
                                                    });
                                                    if(response.status === 200) {
                                                        alert('Blog post deleted successfully!');
                                                        navigate('/');
                                                    }
                                                } catch (error) {
                                                    console.error('Error deleting blog post:', error);
                                                    if (error.response && error.response.status === 403) {
                                                        alert('You do not have permission to delete this blog post.');
                                                    } else {
                                                        alert('Failed to delete blog post. Please try again.');
                                                    }
                                                }
                                            }
                                        }}
                                    >
                                        <i className="fas fa-trash-alt me-1"></i>
                                        Delete
                                    </button>
                                </div>
                            )}
                        </div>
                    </div>
                </Col>
            </Row>
        </Container>
    );
};

export default Blog;