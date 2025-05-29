import React, { useState, useEffect } from 'react';
import { Container, Row, Col, Card, Form, Button, Alert, Tab, Nav, Image } from 'react-bootstrap';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import { useAuth } from '../context/AuthContext';
import { useTheme } from '../context/ThemeContext';
import '../styles/Profile.css';

const Profile = () => {
  const { user, token, isAuthenticated, login, logout } = useAuth();
  const { theme } = useTheme();
  const navigate = useNavigate();

  // 用户资料状态
  const [profileData, setProfileData] = useState({
    username: user?.username || '',
    email: user?.email || '',
    avatar: user?.avatar || '',
    bio: user?.bio || ''
  });

  // 用户发布的文章
  const [userPosts, setUserPosts] = useState([]);

  // 加载状态
  const [loading, setLoading] = useState(false);
  const [postsLoading, setPostsLoading] = useState(false);
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(null);

  // 头像上传状态
  const [avatarFile, setAvatarFile] = useState(null);
  const [avatarPreview, setAvatarPreview] = useState(null);
  const [uploadProgress, setUploadProgress] = useState(0);

  // 密码修改状态
  const [passwordData, setPasswordData] = useState({
    currentPassword: '',
    newPassword: '',
    confirmPassword: ''
  });

  // 在组件加载时获取用户发布的文章
  useEffect(() => {
    if (isAuthenticated && user?.id) {
      fetchUserPosts();
    } else {
      // 如果用户未登录，重定向到登录页面
      navigate('/login', { state: { from: '/profile', message: '请先登录以访问个人资料页面' } });
    }
  }, [isAuthenticated, user, navigate]);

  // 获取用户发布的文章
  const fetchUserPosts = async () => {
    setPostsLoading(true);
    try {
      const response = await axios.get(`${process.env.REACT_APP_API_ROOT}/api/user/posts`, {
        headers: { Authorization: `Bearer ${token}` }
      });

      if (response.status === 200) {
        setUserPosts(response.data.posts || []);
      }
    } catch (err) {
      console.error('获取用户文章失败:', err);
    } finally {
      setPostsLoading(false);
    }
  };

  // 处理表单输入变化
  const handleProfileChange = (e) => {
    const { name, value } = e.target;
    setProfileData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  // 处理密码表单输入变化
  const handlePasswordChange = (e) => {
    const { name, value } = e.target;
    setPasswordData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  // 处理头像文件选择
  const handleAvatarChange = (e) => {
    const file = e.target.files[0];
    if (file) {
      setAvatarFile(file);

      // 创建预览URL
      const previewUrl = URL.createObjectURL(file);
      setAvatarPreview(previewUrl);
    }
  };

  // 上传头像
  const handleAvatarUpload = async () => {
    if (!avatarFile) return;

    setLoading(true);
    setError(null);
    setSuccess(null);

    const formData = new FormData();
    formData.append('avatar', avatarFile);

    try {
      console.log('开始上传头像...');
      const response = await axios.post(
        `${process.env.REACT_APP_API_ROOT}/api/user/avatar`,
        formData,
        {
          headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${token}`
          },
          onUploadProgress: (progressEvent) => {
            const percentCompleted = Math.round((progressEvent.loaded * 100) / progressEvent.total);
            setUploadProgress(percentCompleted);
          }
        }
      );

      console.log('头像上传响应:', response.data);

      if (response.status === 200) {
        setSuccess('头像上传成功！');

        // 获取完整的头像URL
        const avatarUrl = response.data.avatarUrl;
        console.log('头像URL:', avatarUrl);
        console.log('头像上传响应数据:', response.data);

        // 检查是否返回了新令牌
        if (response.data.token) {
          // 使用服务器返回的用户信息
          const updatedUser = response.data.user;
          console.log('头像上传成功，使用服务器返回的用户信息:', updatedUser);

          // 确保头像URL正确
          if (updatedUser.avatar && !updatedUser.avatar.startsWith('http')) {
            console.log('修正头像URL:', updatedUser.avatar);
          }

          // 更新本地存储和认证上下文
          localStorage.setItem('token', response.data.token);
          localStorage.setItem('user', JSON.stringify(updatedUser));
          login(updatedUser, response.data.token);

          // 强制刷新页面以确保头像显示正确
          window.location.reload();
        } else {
          // 如果没有返回新令牌，使用旧方法更新
          const updatedUser = {
            ...user,
            avatar: avatarUrl
          };
          console.log('头像上传成功，手动更新用户信息:', updatedUser);

          // 更新本地存储和认证上下文
          localStorage.setItem('user', JSON.stringify(updatedUser));
          login(updatedUser, token);

          // 强制刷新页面以确保头像显示正确
          window.location.reload();
        }

        // 清除预览
        setAvatarFile(null);
        // 清除预览URL
        if (avatarPreview) {
          URL.revokeObjectURL(avatarPreview);
          setAvatarPreview(null);
        }

        // 更新表单数据
        setProfileData(prev => ({
          ...prev,
          avatar: avatarUrl
        }));
      }
    } catch (err) {
      console.error('头像上传失败:', err);
      setError(err.response?.data?.message || '头像上传失败，请稍后重试');
    } finally {
      setLoading(false);
      setUploadProgress(0);
    }
  };

  // 更新用户资料
  const handleProfileUpdate = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    setSuccess(null);

    try {
      const response = await axios.put(
        `${process.env.REACT_APP_API_ROOT}/api/user/profile`,
        profileData,
        {
          headers: { Authorization: `Bearer ${token}` }
        }
      );

      if (response.status === 200) {
        setSuccess('个人资料更新成功！');

        // 更新用户信息
        const updatedUser = {
          ...user,
          ...profileData
        };

        // 更新本地存储和认证上下文
        localStorage.setItem('user', JSON.stringify(updatedUser));
        login(updatedUser, token);
      }
    } catch (err) {
      console.error('更新个人资料失败:', err);
      setError(err.response?.data?.message || '更新个人资料失败，请稍后重试');
    } finally {
      setLoading(false);
    }
  };

  // 更新密码
  const handlePasswordUpdate = async (e) => {
    e.preventDefault();

    // 验证新密码和确认密码是否匹配
    if (passwordData.newPassword !== passwordData.confirmPassword) {
      setError('新密码和确认密码不匹配');
      return;
    }

    setLoading(true);
    setError(null);
    setSuccess(null);

    try {
      const response = await axios.put(
        `${process.env.REACT_APP_API_ROOT}/api/user/password`,
        passwordData,
        {
          headers: { Authorization: `Bearer ${token}` }
        }
      );

      if (response.status === 200) {
        setSuccess('密码更新成功！');

        console.log('密码更新成功，获取新令牌:', response.data);

        // 更新令牌和用户信息
        if (response.data.token) {
          // 更新用户信息
          const updatedUser = response.data.user;

          // 更新本地存储和认证上下文
          localStorage.setItem('token', response.data.token);
          localStorage.setItem('user', JSON.stringify(updatedUser));
          login(updatedUser, response.data.token);

          // 显示成功消息，但不立即刷新页面
          setTimeout(() => {
            alert('密码已成功更新，请使用新密码重新登录。');
            // 登出并重定向到登录页面
            logout();
            navigate('/login');
          }, 1500);
        }

        // 清空密码表单
        setPasswordData({
          currentPassword: '',
          newPassword: '',
          confirmPassword: ''
        });
      }
    } catch (err) {
      console.error('更新密码失败:', err);
      setError(err.response?.data?.message || '更新密码失败，请稍后重试');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Container className="py-5">
      <Row>
        <Col lg={3} md={4} className="mb-4">
          <Card className={`profile-sidebar ${theme === 'dark' ? 'dark-card' : ''}`}>
            <Card.Body className="text-center">
              <div className="avatar-container mb-4">
                <Image
                  src={
                    avatarPreview ||
                    (profileData.avatar
                      ? (profileData.avatar.startsWith('http')
                          ? profileData.avatar
                          : `${process.env.REACT_APP_API_ROOT}${profileData.avatar}`)
                      : `https://ui-avatars.com/api/?name=${user?.username || 'User'}&background=random&color=fff&size=128`)
                  }
                  onError={(e) => {
                    console.log('头像加载失败，使用默认头像');
                    e.target.onerror = null;
                    e.target.src = `https://ui-avatars.com/api/?name=${user?.username || 'User'}&background=random&color=fff&size=128`;
                  }}
                  roundedCircle
                  className="profile-avatar"
                  alt={user?.username || 'User'}
                />
                <div className="avatar-overlay">
                  <label htmlFor="avatar-upload" className="avatar-edit-btn">
                    <i className="fas fa-camera"></i>
                  </label>
                  <input
                    type="file"
                    id="avatar-upload"
                    className="d-none"
                    accept="image/*"
                    onChange={handleAvatarChange}
                  />
                </div>
              </div>

              {avatarFile && (
                <div className="mb-3">
                  <Button
                    variant="primary"
                    className="w-100 mb-2"
                    onClick={handleAvatarUpload}
                    disabled={loading}
                  >
                    {loading ? '上传中...' : '上传头像'}
                  </Button>

                  {uploadProgress > 0 && (
                    <div className="progress">
                      <div
                        className="progress-bar"
                        role="progressbar"
                        style={{ width: `${uploadProgress}%` }}
                        aria-valuenow={uploadProgress}
                        aria-valuemin="0"
                        aria-valuemax="100"
                      >
                        {uploadProgress}%
                      </div>
                    </div>
                  )}
                </div>
              )}

              <h4 className="profile-username">{user?.username || 'User'}</h4>
              <p className="profile-email text-muted">{user?.email || ''}</p>

              <div className="profile-stats">
                <div className="stat-item">
                  <div className="stat-value">{userPosts.length}</div>
                  <div className="stat-label">文章</div>
                </div>
              </div>
            </Card.Body>
          </Card>
        </Col>

        <Col lg={9} md={8}>
          <Card className={`profile-content ${theme === 'dark' ? 'dark-card' : ''}`}>
            <Card.Body>
              <Tab.Container defaultActiveKey="posts">
                <Nav variant="tabs" className="profile-tabs mb-4">
                  <Nav.Item>
                    <Nav.Link eventKey="posts">My blog</Nav.Link>
                  </Nav.Item>
                  <Nav.Item>
                    <Nav.Link eventKey="edit">Edit profile</Nav.Link>
                  </Nav.Item>
                  <Nav.Item>
                    <Nav.Link eventKey="password">Set new password</Nav.Link>
                  </Nav.Item>
                </Nav>

                <Tab.Content>
                  <Tab.Pane eventKey="posts">
                    {postsLoading ? (
                      <div className="text-center py-5">
                        <div className="spinner-border text-primary" role="status">
                          <span className="visually-hidden">Loading...</span>
                        </div>
                      </div>
                    ) : userPosts.length > 0 ? (
                      <div className="user-posts">
                        {userPosts.map(post => (
                          <div key={post.id} className="post-item">
                            <h5 className="post-title">{post.Title}</h5>
                            <p className="post-excerpt">{post.Post.substring(0, 100)}...</p>
                            <div className="post-actions">
                              <Button
                                variant="outline-primary"
                                size="sm"
                                onClick={() => navigate(`/blog/${post.id}`)}
                              >
                                Look
                              </Button>
                              <Button
                                variant="outline-secondary"
                                size="sm"
                                onClick={() => navigate(`/edit/${post.id}`)}
                              >
                                Edit
                              </Button>
                            </div>
                          </div>
                        ))}
                      </div>
                    ) : (
                      <div className="text-center py-5">
                        <p className="mb-3">You haven't published any articles yet</p>
                        <Button
                          variant="primary"
                          onClick={() => navigate('/add')}
                        >
                          写新文章
                        </Button>
                      </div>
                    )}
                  </Tab.Pane>

                  <Tab.Pane eventKey="edit">
                    {error && <Alert variant="danger">{error}</Alert>}
                    {success && <Alert variant="success">{success}</Alert>}

                    <Form onSubmit={handleProfileUpdate}>
                      <Form.Group className="mb-3">
                        <Form.Label>用户名</Form.Label>
                        <Form.Control
                          type="text"
                          name="username"
                          value={profileData.username}
                          onChange={handleProfileChange}
                          disabled
                        />
                        <Form.Text className="text-muted">
                          用户名不可更改
                        </Form.Text>
                      </Form.Group>

                      <Form.Group className="mb-3">
                        <Form.Label>电子邮箱</Form.Label>
                        <Form.Control
                          type="email"
                          name="email"
                          value={profileData.email}
                          onChange={handleProfileChange}
                          disabled
                        />
                        <Form.Text className="text-muted">
                          邮箱不可更改
                        </Form.Text>
                      </Form.Group>

                      <Form.Group className="mb-3">
                        <Form.Label>个人简介</Form.Label>
                        <Form.Control
                          as="textarea"
                          rows={4}
                          name="bio"
                          value={profileData.bio}
                          onChange={handleProfileChange}
                          placeholder="介绍一下自己吧..."
                        />
                      </Form.Group>

                      <Button
                        variant="primary"
                        type="submit"
                        disabled={loading}
                      >
                        {loading ? '保存中...' : '保存更改'}
                      </Button>
                    </Form>
                  </Tab.Pane>

                  <Tab.Pane eventKey="password">
                    {error && <Alert variant="danger">{error}</Alert>}
                    {success && <Alert variant="success">{success}</Alert>}

                    <Form onSubmit={handlePasswordUpdate}>
                      <Form.Group className="mb-3">
                        <Form.Label>当前密码</Form.Label>
                        <Form.Control
                          type="password"
                          name="currentPassword"
                          value={passwordData.currentPassword}
                          onChange={handlePasswordChange}
                          required
                        />
                      </Form.Group>

                      <Form.Group className="mb-3">
                        <Form.Label>新密码</Form.Label>
                        <Form.Control
                          type="password"
                          name="newPassword"
                          value={passwordData.newPassword}
                          onChange={handlePasswordChange}
                          required
                        />
                      </Form.Group>

                      <Form.Group className="mb-3">
                        <Form.Label>确认新密码</Form.Label>
                        <Form.Control
                          type="password"
                          name="confirmPassword"
                          value={passwordData.confirmPassword}
                          onChange={handlePasswordChange}
                          required
                        />
                      </Form.Group>

                      <Button
                        variant="primary"
                        type="submit"
                        disabled={loading}
                      >
                        {loading ? '更新中...' : '更新密码'}
                      </Button>
                    </Form>
                  </Tab.Pane>
                </Tab.Content>
              </Tab.Container>
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default Profile;
