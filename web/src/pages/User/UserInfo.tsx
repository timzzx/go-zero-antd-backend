import { PageContainer, ProCard, ProColumns, ProForm, ProFormText, ProTable } from '@ant-design/pro-components';
import { message } from 'antd';
import { editPassword } from '@/services/user/api';
import React from 'react';
import { history } from '@umijs/max';

const UserInfo: React.FC = () => {
    return (
        <PageContainer>
            <ProCard title="用户信息编辑">
                <ProForm
                    onFinish={async (values: any) => {
                        const data = await editPassword(values);
                        if (data.code == 200) {
                            message.success("修改成功")
                            localStorage.setItem("token", '');
                            history.push('/user/login');
                        }
                    }}
                >
                    <ProFormText
                        name="password"
                        label="密码"
                        rules={[
                            {
                                required: true,
                                message: '密码长度在6-20',
                                min: 6,
                                max: 20,
                            },
                        ]}
                    >
                    </ProFormText>
                </ProForm>
            </ProCard>
        </PageContainer>

    )
}

export default UserInfo;