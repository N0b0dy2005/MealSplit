import { gql } from '@apollo/client/core';
import apolloClient from './apolloClient.ts';



export interface Activities {
	id: number;
	type: string;
	description: string;
	fromUserId: number;
	toUserId: number;
	userId: number;
	amount: string;
	date: string;
}

export interface Debt {
	id: number;
	fromUserId: number;
	toUserId: number;
	amount: string;
	mealsCount: number;
	createDate: string;
	updatedDate: string;
}

export interface Meal {
	id: number;
	name: string;
	date: string;
	totalAmount: string;
	userId: number;
	description: string;
	createDate: string;
	updatedDate: string;
	userIds: string;
}

export interface MealInput {
	name: string;
	date: string;
	userId: number;
	totalAmount: string;
	userIds: string;
	description: string;
}

export interface Payment {
	id: number;
	fromUserId: number;
	toUserId: number;
	amount: string;
	description: string;
	date: string;
}

export interface PaymentInput {
	fromUserId: number;
	toUserId: number;
	amount: string;
	description: string;
}

export interface User {
	id: number;
	name: string;
	email: string;
	debts: string;
	credits: string;
	createDate: string;
	updatedDate: string;
	phoneNumber: string;
}

export interface response {
	status: string;
	message: string;
	success: boolean;
}

async function createMeal(meal: MealInput): Promise<[response | null, any | null]> {
	try {
		const response = await apolloClient.mutate({
			mutation: gql(`
				 mutation createMeal($meal: MealInput!) {
					createMeal(meal: $meal) {
						status 
						message 
						success 
					}
				}
				`),
			variables: {
								meal: meal
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.createMeal) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.createMeal, null];
	} catch (error) {
		return [null, error];
	}
}

async function createPayment(payment: PaymentInput): Promise<[response | null, any | null]> {
	try {
		const response = await apolloClient.mutate({
			mutation: gql(`
				 mutation createPayment($payment: PaymentInput!) {
					createPayment(payment: $payment) {
						status 
						message 
						success 
					}
				}
				`),
			variables: {
								payment: payment
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.createPayment) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.createPayment, null];
	} catch (error) {
		return [null, error];
	}
}

async function createPost(content: string): Promise<[response | null, any | null]> {
	try {
		const response = await apolloClient.mutate({
			mutation: gql(`
				 mutation createPost($content: String!) {
					createPost(content: $content) {
						status 
						message 
						success 
					}
				}
				`),
			variables: {
								content: content
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.createPost) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.createPost, null];
	} catch (error) {
		return [null, error];
	}
}

async function createUser(name: string, email: string): Promise<[response | null, any | null]> {
	try {
		const response = await apolloClient.mutate({
			mutation: gql(`
				 mutation createUser($name: String!, $email: String!) {
					createUser(name: $name, email: $email) {
						status 
						message 
						success 
					}
				}
				`),
			variables: {
								name: name,
				email: email
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.createUser) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.createUser, null];
	} catch (error) {
		return [null, error];
	}
}

async function getActivities(): Promise<[Activities[] | null, any | null]> {
	try {
		const response = await apolloClient.query({
			query: gql(`
				 query getActivities {
					getActivities {
						id 
						type 
						description 
						fromUserId 
						toUserId 
						userId 
						amount 
						date 
					}
				}
				`),
			variables: {
				
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.getActivities) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.getActivities, null];
	} catch (error) {
		return [null, error];
	}
}

async function getDebts(): Promise<[Debt[] | null, any | null]> {
	try {
		const response = await apolloClient.query({
			query: gql(`
				 query getDebts {
					getDebts {
						id 
						fromUserId 
						toUserId 
						amount 
						mealsCount 
						createDate 
						updatedDate 
					}
				}
				`),
			variables: {
				
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.getDebts) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.getDebts, null];
	} catch (error) {
		return [null, error];
	}
}

async function getMeals(): Promise<[Meal[] | null, any | null]> {
	try {
		const response = await apolloClient.query({
			query: gql(`
				 query getMeals {
					getMeals {
						id 
						name 
						date 
						totalAmount 
						userId 
						description 
						createDate 
						updatedDate 
						userIds 
					}
				}
				`),
			variables: {
				
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.getMeals) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.getMeals, null];
	} catch (error) {
		return [null, error];
	}
}

async function getPayments(): Promise<[Payment[] | null, any | null]> {
	try {
		const response = await apolloClient.query({
			query: gql(`
				 query getPayments {
					getPayments {
						id 
						fromUserId 
						toUserId 
						amount 
						description 
						date 
					}
				}
				`),
			variables: {
				
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.getPayments) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.getPayments, null];
	} catch (error) {
		return [null, error];
	}
}

async function getUsers(): Promise<[User[] | null, any | null]> {
	try {
		const response = await apolloClient.query({
			query: gql(`
				 query getUsers {
					getUsers {
						id 
						name 
						email 
						debts 
						credits 
						createDate 
						updatedDate 
						phoneNumber 
					}
				}
				`),
			variables: {
				
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.getUsers) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.getUsers, null];
	} catch (error) {
		return [null, error];
	}
}

export const query = {
	getActivities,
	getDebts,
	getMeals,
	getPayments,
	getUsers
};
export const mutation = {
	createMeal,
	createPayment,
	createPost,
	createUser
};


