import React from 'react';
import { Link } from 'react-router-dom';
import { Card, Badge } from 'react-bootstrap';
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
import './BlogCard.css';

const BlogCard = ({ blog, onEdit, onDelete, onReadMore }) => {
  // 计算阅读时间（假设每分钟阅读200个字）
  const calculateReadTime = (content) => {
    const words = content.trim().split(/\s+/).length;
    const readTime = Math.ceil(words / 200);
    return readTime < 1 ? 1 : readTime;
  };

  // 格式化日期
  const formatDate = (dateString) => {
    if (!dateString) return '';
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric'
    });
  };

  // 创建日期对象（如果博客没有日期字段，使用当前日期）
  const createdAt = blog.CreatedAt ? new Date(blog.CreatedAt) : new Date();
  const formattedDate = formatDate(createdAt);
  const readTime = calculateReadTime(blog.Post);

  return (
    <Card className="blog-card">
      {/* 博客图片（如果有图片显示图片，如果没有显示占位图） */}
      <div className={`blog-card-image-container ${!blog.Image ? 'no-image' : ''}`}>
        {blog.Image ? (
          <Card.Img
            variant="top"
            src={blog.Image.startsWith('http') ? blog.Image : `${process.env.REACT_APP_API_ROOT}${blog.Image}`}
            className="blog-card-image"
            onError={(e) => {
              e.target.onerror = null;
              e.target.src = 'https://via.placeholder.com/800x400?text=Image+Not+Available';
            }}
          />
        ) : (
          <div className="blog-card-placeholder">
            <div className="placeholder-icon">
              <i className="fas fa-book-open"></i>
            </div>
          </div>
        )}
      </div>

      <Card.Body className="blog-card-body">
        {/* 作者信息和日期 */}
        <div className="blog-card-meta">
          <div className="blog-card-author">
            <img
              src={`https://ui-avatars.com/api/?name=${blog.Author || 'Anonymous'}&background=random&color=fff&size=32`}
              alt={blog.Author || 'Anonymous'}
              className="author-avatar"
            />
            <span>{blog.Author || 'Anonymous'}</span>
          </div>
          <div className="blog-card-date">
            <i className="far fa-calendar-alt"></i>
            <span>{formattedDate}</span>
          </div>
        </div>

        {/* 博客标题 */}
        <Card.Title className="blog-card-title">
          <Link to={`/blog/${blog.id}`}>{blog.Title}</Link>
        </Card.Title>

        {/* 博客摘要 */}
        <Card.Text className="blog-card-excerpt">
          <ReactMarkdown
            remarkPlugins={[remarkGfm]}
            components={{
              // 不显示图片
              img: () => null,
              // 简化链接
              a: ({children}) => <span>{children}</span>,
              // 简化代码块
              code: () => {
                // 内联代码和代码块都简化为文本
                return <span className="code-placeholder">[code]</span>;
              },
              // 简化段落
              p: ({children}) => {
                // 将内容转换为纯文本
                const textContent = React.Children.toArray(children)
                  .map(child => {
                    if (typeof child === 'string') return child;
                    // 如果是React元素且是代码块，返回[code]标记
                    if (React.isValidElement(child) && child.type === 'code') {
                      return '[code]';
                    }
                    return '';
                  })
                  .join('');

                // 截断文本（固定3行，CSS控制）
                return <p>{textContent}</p>;
              },
              // 简化预格式化文本
              pre: () => {
                return <span className="code-placeholder">[code block]</span>;
              }
            }}
          >
            {blog.Post}
          </ReactMarkdown>
        </Card.Text>

        {/* 阅读更多按钮和元数据 */}
        <div className="blog-card-footer">
          <div className="blog-card-meta-tags">
            <Badge bg="light" text="dark" className="read-time">
              <i className="far fa-clock"></i> {readTime} min read
            </Badge>
          </div>

          <div className="blog-card-actions">
            {onEdit && (
              <button
                className="btn btn-sm btn-outline-secondary me-2"
                onClick={() => onEdit(blog.id)}
              >
                <i className="fas fa-edit"></i>
              </button>
            )}

            {onDelete && (
              <button
                className="btn btn-sm btn-outline-danger me-2"
                onClick={() => onDelete(blog.id)}
              >
                <i className="fas fa-trash-alt"></i>
              </button>
            )}

            <Link
              to={`/blog/${blog.id}`}
              className="btn btn-sm btn-primary read-more-btn"
              onClick={(e) => onReadMore && onReadMore(e, blog.id)}
            >
              read more <i className="fas fa-arrow-right ms-1"></i>
            </Link>
          </div>
        </div>
      </Card.Body>
    </Card>
  );
};

export default BlogCard;
