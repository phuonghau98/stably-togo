import React, { useState } from 'react';
import { Alert, Button, Col, Form, InputNumber, Row } from 'antd';
import './App.css';

function App() {
  const [errorMsg, setErrorMsg] = useState('');
  const [responsePrimeNumber, setResponseNumber] = useState<number | null>();
  const [requestedNumber, setRequestedNumber] = useState<number | null>();
  const [isFetchingResult, setFetchingResult] = useState(false);
  function handleFormSubmmited(values: { num: number }) {
    async function fetch(): Promise<any> {
      return new Promise((resolve, reject) => {
        setTimeout(() => {
          const randomState = Math.floor(Math.random() * 10) + 1;
          console.log(randomState);
          if (randomState % 2 === 0) resolve({ num: randomState });
          else reject({ error: 'Something went wrong' });
        }, 500);
      });
    }

    // restart state
    setResponseNumber(null);
    setRequestedNumber(values.num);
    setFetchingResult(true);
    setErrorMsg('')

    // Fetch response
    fetch()
      .then((data) => {
        if ('num' in data) {
          setResponseNumber(data.num);
          setErrorMsg('');
        }
      })
      .catch((err) => {
        setErrorMsg(err.error);
      })
      .finally(() => {
        setFetchingResult(false);
      });
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
                disabled={isFetchingResult}
                size="large"
                style={{ width: '100%' }}
                step={1}
                parser={(v) => parseInt(v || '0')}
              />
            </Form.Item>
            <Button
              loading={isFetchingResult}
              type="primary"
              htmlType="submit"
              size="large"
              style={{ width: '100%' }}
            >
              {isFetchingResult ? 'Figuring it out...' : 'Find'}
            </Button>
          </Form>
          <div style={{ minHeight: 40, marginTop: 20 }}>
            {errorMsg ? (
              <Alert message={<>{errorMsg}</>} type="error" />
            ) : (
              requestedNumber &&
              responsePrimeNumber &&
              responsePrimeNumber !== -1 && (
                <Alert
                  message={
                    <>
                      The highest prime number lower than {requestedNumber} is:{' '}
                      <b>{responsePrimeNumber}</b>
                    </>
                  }
                  type="success"
                />
              )
            )}
          </div>
        </Col>
      </Row>
    </div>
  );
}

export default App;
