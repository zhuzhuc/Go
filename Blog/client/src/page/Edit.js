import axios from "axios";
import React, { useState, useEffect, useRef } from "react";
import { Container, Spinner, Row, Col, Button, Alert } from "react-bootstrap";
import MDEditor, { commands } from '@uiw/react-md-editor';
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';

import { useForm } from "react-hook-form";
import { useNavigate, useParams } from "react-router-dom";
import ImageUploader from "../components/ImageUploader";
import { useAuth } from '../context/AuthContext';

const Edit = () => {
  const [loading, setLoading] = useState(false);
  const [markdownContent, setMarkdownContent] = useState("");
  const [imagePreview, setImagePreview] = useState(null);
  const [previewMode, setPreviewMode] = useState(false);
  const [removeImage, setRemoveImage] = useState(false);
  const [showImageUploader, setShowImageUploader] = useState(false);
  const [permissionError, setPermissionError] = useState(null);
  const editorRef = useRef(null);

  const navigate = useNavigate();
  const { isAuthenticated, user, token } = useAuth();

  // 1. fetch the data
  const params = useParams();
  const [apiData, setApiData] = useState(false);

  // 检查当前用户是否是博客作者
  const isAuthor = () => {
    if (!isAuthenticated || !user || !apiData) return false;
    return apiData.AuthorID === user.id;
  };

  // 检查用户是否已登录
  useEffect(() => {
    if (!isAuthenticated) {
      setPermissionError('您需要登录才能编辑博客');
      navigate('/login', { state: { from: `/edit/${params.id}`, message: '您需要登录才能编辑博客' } });
    }
  }, [isAuthenticated, navigate, params.id]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const apiUrl = process.env.REACT_APP_API_ROOT + "/" + params.id;
        const response = await axios.get(apiUrl);

        if (response.status === 200) {
          console.log("Response data:", response.data);
          const record = response?.data?.record;
          setApiData(record);

          // 设置 Markdown 内容
          if (record && record.Post) {
            setMarkdownContent(record.Post);
          }

          // 检查当前用户是否是博客作者
          if (isAuthenticated && user && record && record.AuthorID !== 0 && record.AuthorID !== user.id) {
            setPermissionError('您没有权限编辑这篇博客');
            navigate('/', { state: { error: '您没有权限编辑这篇博客' } });
          }
        }
      } catch (error) {
        console.log(error.response);
        if (error.response && error.response.status === 404) {
          setPermissionError('找不到博客');
        }
      }
    };

    if (isAuthenticated) {
      fetchData();
    }

    return () => {};
  }, [params.id, isAuthenticated, user, navigate]);

  // 2. Form handling and saving
  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
    watch
  } = useForm();

  // 监听图片上传
  const imageFile = watch("image");

  // 当选择图片时生成预览
  useEffect(() => {
    if (imageFile && imageFile.length > 0) {
      const file = imageFile[0];
      const reader = new FileReader();
      reader.onloadend = () => {
        setImagePreview(reader.result);
      };
      reader.readAsDataURL(file);
    }
  }, [imageFile]);

  // 处理内联图片上传
  const handleImageUploaded = (imagePath) => {
    // 获取当前光标位置
    const textState = editorRef.current?.textareaRef?.getSelectionInfo();
    if (textState) {
      // 在光标位置插入图片 Markdown
      const imageMarkdown = `\n![图片](${imagePath})\n`;
      const newContent =
        markdownContent.substring(0, textState.selection.start) +
        imageMarkdown +
        markdownContent.substring(textState.selection.end);

      setMarkdownContent(newContent);
    } else {
      // 如果无法获取光标位置，则在末尾添加
      const imageMarkdown = `\n![图片](${imagePath})\n`;
      setMarkdownContent(markdownContent + imageMarkdown);
    }
  };

  // 自定义图片上传命令
  const imageUploadCommand = {
    name: 'image-upload',
    keyCommand: 'image-upload',
    buttonProps: { 'aria-label': '上传图片' },
    icon: (
      <svg width="12" height="12" viewBox="0 0 20 20">
        <path fill="currentColor" d="M15 9c1.1 0 2-.9 2-2s-.9-2-2-2-2 .9-2 2 .9 2 2 2zm4-7H1c-.55 0-1 .45-1 1v14c0 .55.45 1 1 1h18c.55 0 1-.45 1-1V3c0-.55-.45-1-1-1zm-1 13l-6-5-2 2-4-5-4 8V4h16v11z" />
      </svg>
    ),
    execute: () => {
      setShowImageUploader(true);
    },
  };

  const saveForm = async (data) => {
    setLoading(true);
    console.log("Form data before sending:", data);

    try {
      const apiUrl = process.env.REACT_APP_API_ROOT + "/" + params.id;

      // 使用 FormData 处理文件上传
      const formData = new FormData();
      formData.append('Title', data.title);

      // 使用 Markdown 内容而不是普通文本
      formData.append('Post', markdownContent);

      // 如果要删除图片
      if (removeImage) {
        formData.append('removeImage', 'true');
      }
      // 如果有选择文件，则添加到 FormData
      else if (data.image && data.image.length > 0) {
        formData.append('file', data.image[0]);
        console.log("Adding file to form data:", data.image[0]);
      }

      // 发送 FormData 而不是 JSON
      const response = await axios.put(apiUrl, formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
          'Authorization': `Bearer ${token || localStorage.getItem('token')}`
        }
      });

      if (response.status === 200) {
        console.log("Blog updated successfully:", response);
        navigate("/", {
          state: "Saved successfully",
        });
      }

      setLoading(false);
    } catch (error) {
      setLoading(false);
      console.log("Error updating blog:", error.response);

      // 处理权限错误
      if (error.response && error.response.status === 403) {
        setPermissionError('您没有权限编辑这篇博客');
        alert('您没有权限编辑这篇博客');
        navigate('/');
      } else if (error.response && error.response.status === 401) {
        setPermissionError('您需要登录才能编辑博客');
        alert('您需要登录才能编辑博客');
        navigate('/login', { state: { from: `/edit/${params.id}`, message: '您需要登录才能编辑博客' } });
      } else {
        alert('更新博客失败: ' + (error.response?.data?.message || '未知错误'));
      }
    }
  };

  if (loading) {
    return (
      <>
        <Container className="spinner">
          <Spinner animation="grow" />
        </Container>
      </>
    );
  }

  return (
    <>
      {/* 顶部导航栏 */}
      <div className="bg-gradient" style={{
        background: 'linear-gradient(90deg, #800080, #e91e63)',
        color: 'white',
        padding: '15px 0',
        marginBottom: '30px'
      }}>
        <Container>
          <div className="d-flex justify-content-between align-items-center">
            <h2 className="m-0">PERSONAL BLOG</h2>
            {isAuthenticated ? (
              <div className="d-flex align-items-center">
                <span className="me-3">
                  <i className="fas fa-user me-1"></i> {user?.username || 'User'}
                </span>
                <button
                  className="btn btn-outline-light btn-sm"
                  onClick={() => {
                    localStorage.removeItem('token');
                    localStorage.removeItem('user');
                    navigate('/');
                    window.location.reload();
                  }}
                >
                  <i className="fas fa-sign-out-alt me-1"></i> Logout
                </button>
              </div>
            ) : (
              <button
                className="btn btn-outline-light btn-sm"
                onClick={() => navigate('/login')}
              >
                <i className="fas fa-sign-in-alt me-1"></i> Login
              </button>
            )}
          </div>
        </Container>
      </div>

      <Container>
        {/* 权限错误提示 */}
        {permissionError && (
          <div className="alert alert-danger mb-4" role="alert">
            <i className="fas fa-exclamation-circle me-2"></i>
            {permissionError}
          </div>
        )}

        {/* 导航菜单 */}
        <div className="text-center mb-4">
          <ul className="nav nav-pills justify-content-center">
            <li className="nav-item">
              <button className="nav-link" onClick={() => navigate("/")}>
                <i className="fas fa-home me-1"></i> HOME
              </button>
            </li>
            <li className="nav-item">
              <button className="nav-link active">
                <i className="fas fa-blog me-1"></i> BLOG
              </button>
            </li>
            <li className="nav-item">
              <button className="nav-link">
                <i className="fas fa-info-circle me-1"></i> ABOUT
              </button>
            </li>
            <li className="nav-item">
              <button className="nav-link">
                <i className="fas fa-envelope me-1"></i> CONTACT
              </button>
            </li>
          </ul>
        </div>

        {/* 编辑表单 */}
        <div className="bg-white p-4 rounded shadow-sm" style={{ maxWidth: '800px', margin: '0 auto' }}>
          <h2 className="text-center mb-4">Edit Post</h2>

          {apiData && (
            <form onSubmit={handleSubmit(saveForm)}>
              <div className="mb-4">
                <input
                  type="text"
                  defaultValue={apiData.Title}
                  className={`form-control form-control-lg border-0 shadow-sm ${errors.title ? "is-invalid" : ""}`}
                  placeholder="Post Title"
                  style={{ fontSize: '1.5rem' }}
                  {...register("title", {
                    required: { value: true, message: "Title is required." },
                    minLength: {
                      value: 3,
                      message: "Title should be minimum 3 characters.",
                    },
                  })}
                />
                {errors.title && (
                  <div className="invalid-feedback">{errors.title.message}</div>
                )}
              </div>

              {/* 编辑/预览切换按钮 */}
              <div className="mb-4 d-flex justify-content-end">
                <div className="btn-group">
                  <button
                    type="button"
                    className={`btn ${!previewMode ? 'btn-primary' : 'btn-outline-primary'}`}
                    onClick={() => setPreviewMode(false)}
                  >
                    <i className="fas fa-edit me-1"></i> 编辑
                  </button>
                  <button
                    type="button"
                    className={`btn ${previewMode ? 'btn-primary' : 'btn-outline-primary'}`}
                    onClick={() => setPreviewMode(true)}
                  >
                    <i className="fas fa-eye me-1"></i> 预览
                  </button>
                </div>
              </div>

              {/* Markdown 编辑器或预览 */}
              <div className="mb-4">
                {!previewMode ? (
                  <>
                    <label className="form-label mb-2">博客内容 (支持 Markdown 格式)</label>
                    <div data-color-mode="light">
                      <MDEditor
                        ref={editorRef}
                        value={markdownContent}
                        onChange={setMarkdownContent}
                        height={400}
                        preview="edit"
                        highlightEnable={true}
                        style={{
                          borderRadius: '0.375rem',
                          boxShadow: '0 .125rem .25rem rgba(0,0,0,.075)'
                        }}
                        commands={[
                          ...commands.getCommands(),
                          imageUploadCommand
                        ]}
                      />
                    </div>

                    {/* 图片上传模态框 */}
                    <ImageUploader
                      show={showImageUploader}
                      onHide={() => setShowImageUploader(false)}
                      onImageUploaded={handleImageUploaded}
                      authToken={token}
                    />
                    {!markdownContent && (
                      <div className="text-danger mt-2 small">Post Content is required.</div>
                    )}

                    <div className="mt-3">
                      <h5>Markdown 提示:</h5>
                      <ul className="small text-muted">
                        <li>使用 # 创建标题，例如：# 标题1</li>
                        <li>使用 **文本** 创建粗体文本</li>
                        <li>使用 *文本* 创建斜体文本</li>
                        <li>使用 ![描述](图片URL) 插入图片</li>
                        <li>使用 [链接文本](URL) 创建链接</li>
                      </ul>
                    </div>
                  </>
                ) : (
                  <div className="preview-container">
                    <h4 className="mb-3">文章预览</h4>
                    <div className="border rounded p-4 bg-light">
                      <h2 className="preview-title mb-4">{watch("title") || apiData.Title}</h2>

                      {/* 显示特色图片预览 */}
                      {imagePreview ? (
                        <div className="featured-image mb-4">
                          <img
                            src={imagePreview}
                            alt="Featured"
                            className="img-fluid rounded shadow-sm"
                            style={{ maxHeight: '300px', width: 'auto', display: 'block', margin: '0 auto' }}
                          />
                        </div>
                      ) : apiData.Image && (
                        <div className="featured-image mb-4">
                          <img
                            src={process.env.REACT_APP_API_ROOT + apiData.Image}
                            alt="Featured"
                            className="img-fluid rounded shadow-sm"
                            style={{ maxHeight: '300px', width: 'auto', display: 'block', margin: '0 auto' }}
                            onError={(e) => {
                              console.log("Image load error:", e);
                              console.log("Image URL:", process.env.REACT_APP_API_ROOT + apiData.Image);
                              e.target.onerror = null;
                              e.target.src = 'https://via.placeholder.com/300x200?text=Image+Not+Found';
                            }}
                          />
                        </div>
                      )}

                      {/* 显示 Markdown 内容预览 */}
                      <div className="blog-content">
                        <ReactMarkdown
                          remarkPlugins={[remarkGfm]}
                          components={{
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
                            a: ({node, children, ...props}) => (
                              <a {...props} target="_blank" rel="noopener noreferrer" className="text-decoration-none">{children}</a>
                            ),
                            p: ({node, ...props}) => (
                              <p {...props} className="mb-3" style={{ fontSize: '1.1rem', lineHeight: '1.8' }} />
                            )
                          }}
                        >
                          {markdownContent || "在这里编写您的文章内容..."}
                        </ReactMarkdown>
                      </div>
                    </div>
                  </div>
                )}
              </div>

              {/* 图片上传 */}
              <div className="mb-4">
                <label className="form-label">特色图片</label>
                <div className="input-group mb-3">
                  <span className="input-group-text bg-light">
                    <i className="fas fa-image"></i>
                  </span>
                  <input
                    type="file"
                    className={`form-control ${errors.image ? "is-invalid" : ""}`}
                    accept="image/*"
                    {...register("image")}
                  />
                </div>

                {/* 图片预览 */}
                {imagePreview ? (
                  <div className="image-preview mt-2 mb-3">
                    <div className="d-flex justify-content-between align-items-center mb-2">
                      <p className="text-muted small mb-0">新图片预览:</p>
                      <button
                        type="button"
                        className="btn btn-sm btn-outline-danger"
                        onClick={() => {
                          setImagePreview(null);
                          setValue("image", null);
                        }}
                      >
                        <i className="fas fa-trash-alt me-1"></i> 移除
                      </button>
                    </div>
                    <div className="border rounded p-2 text-center bg-light">
                      <img
                        src={imagePreview}
                        alt="Preview"
                        className="img-fluid"
                        style={{ maxHeight: '200px', width: 'auto' }}
                      />
                    </div>
                  </div>
                ) : apiData.Image && !removeImage ? (
                  <div className="image-preview mt-2 mb-3">
                    <div className="d-flex justify-content-between align-items-center mb-2">
                      <p className="text-muted small mb-0">当前图片:</p>
                      <button
                        type="button"
                        className="btn btn-sm btn-outline-danger"
                        onClick={() => setRemoveImage(true)}
                      >
                        <i className="fas fa-trash-alt me-1"></i> 删除
                      </button>
                    </div>
                    <div className="border rounded p-2 text-center bg-light">
                      <img
                        src={process.env.REACT_APP_API_ROOT + apiData.Image}
                        alt="Current"
                        className="img-fluid"
                        style={{ maxHeight: '200px', width: 'auto' }}
                        onError={(e) => {
                          console.log("Image load error:", e);
                          e.target.onerror = null;
                          e.target.src = 'https://via.placeholder.com/300x200?text=Image+Not+Found';
                        }}
                      />
                    </div>
                  </div>
                ) : removeImage && (
                  <div className="alert alert-warning mt-2 mb-3">
                    <div className="d-flex justify-content-between align-items-center">
                      <span><i className="fas fa-exclamation-triangle me-2"></i> 图片将被删除</span>
                      <button
                        type="button"
                        className="btn btn-sm btn-outline-secondary"
                        onClick={() => setRemoveImage(false)}
                      >
                        取消
                      </button>
                    </div>
                  </div>
                )}

                <div className="form-text text-muted">
                  <small>选择一张图片作为博客文章的特色图片 (可选)</small>
                </div>
              </div>

              <div className="d-flex justify-content-between mt-5">
                <button
                  type="button"
                  className="btn btn-outline-secondary px-4"
                  onClick={() => navigate("/")}
                >
                  <i className="fas fa-times me-2"></i>
                  Cancel
                </button>
                <button type="submit" className="btn px-4 text-white" style={{
                  background: 'linear-gradient(90deg, #800080, #e91e63)'
                }}>
                  <i className="fas fa-save me-2"></i>
                  Save Changes
                </button>
              </div>
            </form>
          )}
        </div>
      </Container>

    </>
  );
};

export default Edit;