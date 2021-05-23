import React from 'react';
import { Button, Col, Form, InputNumber, Row } from 'antd';
import './App.css';

function App() {
  function handleFormSubmmited(values: any) {
    console.log(values);
  }

  return (
    <div className="App">
      <Row justify="center" align="middle" style={{ height: '100%' }}>
        <Col>
          <Form layout="vertical" onFinish={handleFormSubmmited}>
            <Form.Item
              label="Enter a number that you want to find the lower nearest prime"
              name="num"
              initialValue={2}
              rules={[
                {
                  required: true,
                  message: 'A number is required',
                },
              ]}
            >
              <InputNumber
                min={2}
                size="large"
                style={{ width: '100%' }}
                step={1}
                parser={(v) => parseInt(v || '0')}
              />
            </Form.Item>
            <Button
              type="primary"
              htmlType="submit"
              size="large"
              style={{ width: '100%' }}
            >
              Find
            </Button>
          </Form>
        </Col>
      </Row>
    </div>
  );
}

export default App;
