// src/api.js
import axios from 'axios';

const API_URL = process.env.REACT_APP_API_URL;

export const getEmployees = async () => axios.get(`${API_URL}/employees`);
export const getEmployeeById = async (id) => axios.get(`${API_URL}/employees/${id}`);
export const createEmployee = async (employee) => axios.post(`${API_URL}/employees`, employee);
export const updateEmployee = async (id, employee) => axios.put(`${API_URL}/employees/${id}`, employee);
export const deleteEmployee = async (id) => axios.delete(`${API_URL}/employees/${id}`);
