import React from 'react';
import {HashRouter as Router} from 'react-router-dom';
import {Col, Layout, Row, Typography, Alert} from 'antd';
import './App.css';
import CustomMenu from './global/CustomMenu';
import CustomContent from './global/CustomContent';
import UserMenu from './global/UserMenu';
import GlobalProvider from './context/GlobalProvider';

const { Header, Content, Sider } = Layout;

const App = () => (
    <Router>
        <GlobalProvider>
            <Layout>
                <Header className="header">
                    <Row>
                        <Col flex={1}><Typography.Title style={{ color: 'white' }}>Go Shorten!</Typography.Title></Col>
                        <Col>
                            <UserMenu/>
                        </Col>
                    </Row>
                </Header>
                <Layout>
                    <Sider width={200} className="site-layout-background">
                        <Alert.ErrorBoundary>
                            <CustomMenu/>
                        </Alert.ErrorBoundary>
                    </Sider>
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
