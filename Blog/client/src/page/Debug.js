import React, { useState, useEffect } from 'react';
import { Container, Row, Col, Card, Form, Button, Alert, Table } from 'react-bootstrap';
import axios from 'axios';

const Debug = () => {
  const [username, setUsername] = useState('');
  const [newPassword, setNewPassword] = useState('');
  const [avatarUrl, setAvatarUrl] = useState('');
  const [blogTitle, setBlogTitle] = useState('');
  const [blogContent, setBlogContent] = useState('');
  const [blogs, setBlogs] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(null);
  const [userInfo, setUserInfo] = useState(null);

  // 重置密码
  const handleResetPassword = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    setSuccess(null);

    try {
      const response = await axios.post(
        `${process.env.REACT_APP_API_ROOT}/debug/reset-password`,
        { username, newPassword }
      );

      if (response.status === 200) {
        setSuccess('密码重置成功！');
        setNewPassword('');
      }
    } catch (err) {
      console.error('密码重置失败:', err);
      setError(err.response?.data?.message || '密码重置失败，请稍后重试');
    } finally {
      setLoading(false);
    }
  };

  // 获取用户信息
  const handleGetUserInfo = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    setSuccess(null);
    setUserInfo(null);

    try {
      const response = await axios.post(
        `${process.env.REACT_APP_API_ROOT}/debug/user-info`,
        { username }
      );

      if (response.status === 200) {
        setUserInfo(response.data.user);
        setSuccess('获取用户信息成功！');

        // 如果用户有头像，自动填充头像URL
        if (response.data.user.avatar) {
          setAvatarUrl(response.data.user.avatar);
        }
      }
    } catch (err) {
      console.error('获取用户信息失败:', err);
      setError(err.response?.data?.message || '获取用户信息失败，请稍后重试');
    } finally {
      setLoading(false);
    }
  };

  // 更新头像
  const handleUpdateAvatar = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    setSuccess(null);

    try {
      const response = await axios.post(
        `${process.env.REACT_APP_API_ROOT}/debug/update-avatar`,
        { username, avatarUrl }
      );

      if (response.status === 200) {
        setSuccess('头像更新成功！');
        // 刷新用户信息
        handleGetUserInfo(e);
      }
    } catch (err) {
      console.error('头像更新失败:', err);
      setError(err.response?.data?.message || '头像更新失败，请稍后重试');
    } finally {
      setLoading(false);
    }
  };

  // 创建博客
  const handleCreateBlog = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    setSuccess(null);

    try {
      const response = await axios.post(
        `${process.env.REACT_APP_API_ROOT}/debug/blog-create`,
        {
          title: blogTitle,
          content: blogContent,
          username: username
        }
      );

      if (response.status === 200) {
        setSuccess('博客创建成功！');
        setBlogTitle('');
        setBlogContent('');
        // 刷新博客列表
        fetchBlogs();
      }
    } catch (err) {
      console.error('博客创建失败:', err);
      setError(err.response?.data?.message || '博客创建失败，请稍后重试');
    } finally {
      setLoading(false);
    }
  };

  // 获取博客列表
  const fetchBlogs = async () => {
    try {
      const response = await axios.get(
        `${process.env.REACT_APP_API_ROOT}/debug/blog-list`
      );

      if (response.status === 200) {
        setBlogs(response.data.blogs);
      }
    } catch (err) {
      console.error('获取博客列表失败:', err);
    }
  };

  // 组件加载时获取博客列表
  useEffect(() => {
    fetchBlogs();
  }, []);

  return (
    <Container className="py-5">
      <h1 className="text-center mb-4">调试工具</h1>
      {error && <Alert variant="danger">{error}</Alert>}
      {success && <Alert variant="success">{success}</Alert>}

      <Row>
        <Col md={6} className="mb-4">
          <Card className="mb-4">
            <Card.Header>密码重置工具</Card.Header>
            <Card.Body>
              <Form onSubmit={handleResetPassword}>
                <Form.Group className="mb-3">
                  <Form.Label>用户名</Form.Label>
                  <Form.Control
                    type="text"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    required
                  />
                </Form.Group>

                <Form.Group className="mb-3">
                  <Form.Label>新密码</Form.Label>
                  <Form.Control
                    type="password"
                    value={newPassword}
                    onChange={(e) => setNewPassword(e.target.value)}
                    required
                  />
                </Form.Group>

                <Button
                  variant="primary"
                  type="submit"
                  disabled={loading}
                >
                  {loading ? '处理中...' : '重置密码'}
                </Button>
              </Form>
            </Card.Body>
          </Card>

          <Card className="mb-4">
            <Card.Header>头像更新工具</Card.Header>
            <Card.Body>
              <Form onSubmit={handleUpdateAvatar}>
                <Form.Group className="mb-3">
                  <Form.Label>用户名</Form.Label>
                  <Form.Control
                    type="text"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    required
                  />
                </Form.Group>

                <Form.Group className="mb-3">
                  <Form.Label>头像URL</Form.Label>
                  <Form.Control
                    type="text"
                    value={avatarUrl}
                    onChange={(e) => setAvatarUrl(e.target.value)}
                    required
                    placeholder="/uploads/avatars/example.jpg"
                  />
                  <Form.Text className="text-muted">
                    输入相对路径，例如：/uploads/avatars/example.jpg
                  </Form.Text>
                </Form.Group>

                <Button
                  variant="success"
                  type="submit"
                  disabled={loading}
                >
                  {loading ? '更新中...' : '更新头像'}
                </Button>
              </Form>
            </Card.Body>
          </Card>

          <Card>
            <Card.Header>博客创建工具</Card.Header>
            <Card.Body>
              <Form onSubmit={handleCreateBlog}>
                <Form.Group className="mb-3">
                  <Form.Label>用户名</Form.Label>
                  <Form.Control
                    type="text"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    required
                  />
                </Form.Group>

                <Form.Group className="mb-3">
                  <Form.Label>博客标题</Form.Label>
                  <Form.Control
                    type="text"
                    value={blogTitle}
                    onChange={(e) => setBlogTitle(e.target.value)}
                    required
                  />
                </Form.Group>

                <Form.Group className="mb-3">
                  <Form.Label>博客内容</Form.Label>
                  <Form.Control
                    as="textarea"
                    rows={4}
                    value={blogContent}
                    onChange={(e) => setBlogContent(e.target.value)}
                    required
                  />
                </Form.Group>

                <Button
                  variant="primary"
                  type="submit"
                  disabled={loading}
                >
                  {loading ? '创建中...' : '创建博客'}
                </Button>
              </Form>
            </Card.Body>
          </Card>
        </Col>

        <Col md={6}>
          <Card className="mb-4">
            <Card.Header>用户信息查询</Card.Header>
            <Card.Body>
              <Form onSubmit={handleGetUserInfo} className="mb-3">
                <Form.Group className="mb-3">
                  <Form.Label>用户名</Form.Label>
                  <Form.Control
                    type="text"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    required
                  />
                </Form.Group>

                <Button
                  variant="info"
                  type="submit"
                  disabled={loading}
                >
                  {loading ? '查询中...' : '查询用户信息'}
                </Button>
              </Form>

              {userInfo && (
                <div className="mt-4">
                  <h5>用户信息</h5>
                  <pre className="bg-light p-3 rounded">
                    {JSON.stringify(userInfo, null, 2)}
                  </pre>

                  {userInfo.avatar && (
                    <div className="mt-3">
                      <h5>头像预览</h5>
                      <div className="border p-3 text-center">
                        <img
                          src={`${process.env.REACT_APP_API_ROOT}${userInfo.avatar}`}
                          alt="用户头像"
                          style={{ maxWidth: '100%', maxHeight: '200px' }}
                          onError={(e) => {
                            console.log('头像加载失败');
                            e.target.src = `https://ui-avatars.com/api/?name=${userInfo.username}&background=random&color=fff&size=128`;
                          }}
                        />
                      </div>
                    </div>
                  )}
                </div>
              )}
            </Card.Body>
          </Card>

          <Card>
            <Card.Header>博客列表</Card.Header>
            <Card.Body>
              <Button
                variant="outline-primary"
                className="mb-3"
                onClick={fetchBlogs}
              >
                刷新博客列表
              </Button>

              {blogs.length > 0 ? (
                <Table striped bordered hover responsive>
                  <thead>
                    <tr>
                      <th>ID</th>
                      <th>标题</th>
                      <th>作者</th>
                    </tr>
                  </thead>
                  <tbody>
                    {blogs.map(blog => (
                      <tr key={blog.id}>
                        <td>{blog.id}</td>
                        <td>{blog.Title}</td>
                        <td>{blog.Author}</td>
                      </tr>
                    ))}
                  </tbody>
                </Table>
              ) : (
                <p className="text-center">暂无博客</p>
              )}
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default Debug;
