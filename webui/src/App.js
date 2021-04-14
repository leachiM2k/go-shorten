import React from 'react';
import {BrowserRouter as Router} from 'react-router-dom';
import {Col, Layout, Row, Typography, Alert} from 'antd';
import './App.css';
import CustomContent from './global/CustomContent';
import GlobalProvider from './context/GlobalProvider';
import CustomRightHeader from './global/CustomRightHeader';

const { Header, Content, Sider } = Layout;

const App = () => (
    <Router>
        <GlobalProvider>
            <Layout>
                <Header className="header">
                    <Row>
                        <Col flex={1}><Typography.Title style={{ color: 'white' }}>Go Shorten!</Typography.Title></Col>
                        <Col>
                            <CustomRightHeader/>
                        </Col>
                    </Row>
                </Header>
                <Layout>
                    <Layout style={{ padding: '24px' }}>
                        <Content
                            className="site-layout-background"
                            style={{
                                padding: 24,
                                margin: 0,
                                minHeight: 280,
                            }}
                        >
                            <Alert.ErrorBoundary>
                                <CustomContent/>
                            </Alert.ErrorBoundary>
                        </Content>
                    </Layout>
                </Layout>
            </Layout>
        </GlobalProvider>
    </Router>
);

export default App;
