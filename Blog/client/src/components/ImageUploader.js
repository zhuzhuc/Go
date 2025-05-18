import React, { useState } from 'react';
import axios from 'axios';
import { Modal, Button, Spinner } from 'react-bootstrap';
import { useAuth } from '../context/AuthContext';

const ImageUploader = ({ onImageUploaded, show, onHide, authToken }) => {
  const [selectedFile, setSelectedFile] = useState(null);
  const [preview, setPreview] = useState(null);
  const [uploading, setUploading] = useState(false);
  const [error, setError] = useState(null);
  const { token: contextToken } = useAuth();

  // 优先使用传入的 authToken，其次使用 context 中的 token
  const token = authToken || contextToken;

  // 处理文件选择
  const handleFileChange = (event) => {
    const file = event.target.files[0];
    if (file) {
      setSelectedFile(file);

      // 创建预览
      const reader = new FileReader();
      reader.onloadend = () => {
        setPreview(reader.result);
      };
      reader.readAsDataURL(file);

      // 清除错误
      setError(null);
    }
  };

  // 上传图片
  const uploadImage = async () => {
    if (!selectedFile) {
      setError('请选择一个图片文件');
      return;
    }

    setUploading(true);
    setError(null);

    try {
      // 检查 token 是否存在
      console.log("Token in ImageUploader:", token);

      // 如果 token 不存在，尝试从 localStorage 获取
      const authToken = token || localStorage.getItem('token');
      console.log("Using token:", authToken ? `Bearer ${authToken.substring(0, 10)}...` : 'No token');

      if (!authToken) {
        setError('未找到授权令牌，请重新登录');
        setUploading(false);
        return;
      }

      // 打印 localStorage 中的所有内容，帮助调试
      console.log("localStorage contents:");
      for (let i = 0; i < localStorage.length; i++) {
        const key = localStorage.key(i);
        console.log(`${key}: ${key === 'token' ? 'TOKEN_EXISTS' : key === 'user' ? 'USER_DATA_EXISTS' : 'OTHER_DATA'}`);
      }

      const formData = new FormData();
      formData.append('file', selectedFile);

      // 打印完整的请求 URL
      const apiUrl = `${process.env.REACT_APP_API_ROOT}/upload-image`;
      console.log("Upload URL:", apiUrl);

      // 直接从 localStorage 获取 token
      const localStorageToken = localStorage.getItem('token');
      console.log("Token from localStorage:", localStorageToken ? `${localStorageToken.substring(0, 10)}...` : 'No token');

      // 设置请求头，优先使用 localStorage 中的 token
      const headers = {
        'Content-Type': 'multipart/form-data',
        'Authorization': `Bearer ${localStorageToken || authToken}`
      };
      console.log("Request headers:", {
        'Content-Type': headers['Content-Type'],
        'Authorization': headers['Authorization'] ? 'Bearer TOKEN_EXISTS' : 'No Authorization'
      });

      // 发送请求
      const response = await axios.post(
        apiUrl,
        formData,
        { headers }
      );

      if (response.status === 200 || response.status === 201) {
        const imagePath = response.data.path;
        onImageUploaded(imagePath);
        resetForm();
        onHide();
      }
    } catch (err) {
      console.error('上传图片失败:', err);
      setError('上传图片失败: ' + (err.response?.data?.message || err.message));
    } finally {
      setUploading(false);
    }
  };

  // 重置表单
  const resetForm = () => {
    setSelectedFile(null);
    setPreview(null);
    setError(null);
  };

  return (
    <Modal show={show} onHide={onHide} centered>
      <Modal.Header closeButton>
        <Modal.Title>上传图片</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <div className="mb-3">
          <label htmlFor="image-upload" className="form-label">选择图片</label>
          <input
            type="file"
            className="form-control"
            id="image-upload"
            accept="image/*"
            onChange={handleFileChange}
          />
        </div>

        {preview && (
          <div className="image-preview mt-3 mb-3">
            <p className="mb-2 text-muted small">图片预览:</p>
            <div className="border rounded p-2 text-center bg-light">
              <img
                src={preview}
                alt="Preview"
                className="img-fluid"
                style={{ maxHeight: '200px', width: 'auto' }}
              />
            </div>
          </div>
        )}

        {error && (
          <div className="alert alert-danger mt-3" role="alert">
            {error}
          </div>
        )}
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={onHide}>
          取消
        </Button>
        <Button
          variant="primary"
          onClick={uploadImage}
          disabled={!selectedFile || uploading}
        >
          {uploading ? (
            <>
              <Spinner
                as="span"
                animation="border"
                size="sm"
                role="status"
                aria-hidden="true"
                className="me-2"
              />
              上传中...
            </>
          ) : '上传图片'}
        </Button>
      </Modal.Footer>
    </Modal>
  );
};

export default ImageUploader;
