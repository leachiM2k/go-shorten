import {Button, Drawer} from 'antd';
import React from 'react';
import './DrawerForm.css';
import ShortsForm from './ShortsForm';

const DrawerForm = ({ form, title, actionText, onSave, onClose, initialValues, visible }) => {
    const handleSubmit = () => {
        form
            .validateFields()
            .catch(info => {
                console.log('Validate Failed:', info);
            })
            .then(values => onSave(values));
    }

    return (
        <Drawer
            title={title}
            width={""}
            className="mobile-fullscreen-drawer"
            onClose={onClose}
            visible={visible}
            bodyStyle={{ paddingBottom: 80 }}
            destroyOnClose
            footer={
                <div style={{ textAlign: 'right' }}>
                    <Button onClick={onClose} style={{ marginRight: 8 }}>Cancel</Button>
                    <Button onClick={handleSubmit} type="primary">{actionText}</Button>
                </div>
            }
        >
            <ShortsForm form={form} initialValues={initialValues}/>
        </Drawer>
    );
};

export default DrawerForm;
