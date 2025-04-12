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
	isConfirmed: boolean;
}

export interface CurrentMealsResponse {
	id: number;
	name: string;
	totalAmount: string;
	totalMealsCount: number;
}

export interface Dashboard {
	totalUsers: number;
	totalMeals: number;
	totalDebts: string;
	totalCredits: string;
	tobDebtsPerUser: TobDebtsPerUser[];
	totalCreditsPerUser: TotalCreditsPerUser[];
}

export interface Debt {
	id: number;
	fromUserId: number;
	toUserId: number;
	amount: string;
	mealsCount: number;
	createDate: string;
	updatedDate: string;
	isConfirmed: boolean;
	mealId: number;
}

export interface DebtDetail {
	id: number;
	from_user_id: number;
	to_user_id: number;
	from_user_name: string;
	to_user_name: string;
	amount: string;
	meals_count: number;
	create_date: string;
	is_confirmed: boolean;
	type: string;
	meal_id: number;
	payment_id: number;
	meal_info: MealInfo;
	payment_info: PaymentInfo;
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
	products: string;
}

export interface MealInfo {
	id: number;
	name: string;
	description: string;
	date: string;
	total_amount: string;
}

export interface MealInput {
	name: string;
	date: string;
	userId: number;
	totalAmount: string;
	userIds: string;
	description: string;
	produkts: string;
}

export interface Payment {
	id: number;
	fromUserId: number;
	toUserId: number;
	amount: string;
	description: string;
	date: string;
	isConfirmed: boolean;
}

export interface PaymentInfo {
	id: number;
	description: string;
	date: string;
	amount: string;
}

export interface PaymentInput {
	fromUserId: number;
	toUserId: number;
	amount: string;
	description: string;
}

export interface TobDebtsPerUser {
	userId: number;
	amount: string;
}

export interface TotalCreditsPerUser {
	userId: number;
	amount: string;
}

export interface UpdateUserInput {
	id: number;
	name: string;
	email: string;
	phoneNumber: string;
	password: string;
}

export interface Upload {
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
	admin: boolean;
}

export interface UserInput {
	id: number;
	name: string;
	email: string;
	admin: boolean;
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

async function deleteUser(id: number): Promise<[response | null, any | null]> {
	try {
		const response = await apolloClient.mutate({
			mutation: gql(`
				 mutation deleteUser($id: Int!) {
					deleteUser(id: $id) {
						status 
						message 
						success 
					}
				}
				`),
			variables: {
								id: id
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.deleteUser) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.deleteUser, null];
	} catch (error) {
		return [null, error];
	}
}

async function editUser(user: UserInput): Promise<[response | null, any | null]> {
	try {
		const response = await apolloClient.mutate({
			mutation: gql(`
				 mutation editUser($user: UserInput!) {
					editUser(user: $user) {
						status 
						message 
						success 
					}
				}
				`),
			variables: {
								user: user
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.editUser) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.editUser, null];
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
						isConfirmed 
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

async function getCurrentUser(): Promise<[User | null, any | null]> {
	try {
		const response = await apolloClient.query({
			query: gql(`
				 query getCurrentUser {
					getCurrentUser {
						id 
						name 
						email 
						debts 
						credits 
						createDate 
						updatedDate 
						phoneNumber 
						admin 
					}
				}
				`),
			variables: {
				
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.getCurrentUser) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.getCurrentUser, null];
	} catch (error) {
		return [null, error];
	}
}

async function getDashboard(): Promise<[Dashboard | null, any | null]> {
	try {
		const response = await apolloClient.query({
			query: gql(`
				 query getDashboard {
					getDashboard {
						totalUsers 
						totalMeals 
						totalDebts 
						totalCredits 
						tobDebtsPerUser {
							userId 
							amount 
						}
						totalCreditsPerUser {
							userId 
							amount 
						}
					}
				}
				`),
			variables: {
				
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.getDashboard) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.getDashboard, null];
	} catch (error) {
		return [null, error];
	}
}

async function getDebtDetails(fromUserId: number, toUserId: number): Promise<[DebtDetail[] | null, any | null]> {
	try {
		const response = await apolloClient.query({
			query: gql(`
				 query getDebtDetails($fromUserId: Int!, $toUserId: Int!) {
					getDebtDetails(fromUserId: $fromUserId, toUserId: $toUserId) {
						id 
						from_user_id 
						to_user_id 
						from_user_name 
						to_user_name 
						amount 
						meals_count 
						create_date 
						is_confirmed 
						type 
						meal_id 
						payment_id 
						meal_info {
							id 
							name 
							description 
							date 
							total_amount 
						}
						payment_info {
							id 
							description 
							date 
							amount 
						}
					}
				}
				`),
			variables: {
								fromUserId: fromUserId,
				toUserId: toUserId
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.getDebtDetails) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.getDebtDetails, null];
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
						isConfirmed 
						mealId 
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

async function getMealById(id: number): Promise<[Meal | null, any | null]> {
	try {
		const response = await apolloClient.query({
			query: gql(`
				 query getMealById($id: Int!) {
					getMealById(id: $id) {
						id 
						name 
						date 
						totalAmount 
						userId 
						description 
						createDate 
						updatedDate 
						userIds 
						products 
					}
				}
				`),
			variables: {
								id: id
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.getMealById) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.getMealById, null];
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
						products 
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

async function getMealsForCurrentUser(): Promise<[Meal[] | null, any | null]> {
	try {
		const response = await apolloClient.query({
			query: gql(`
				 query getMealsForCurrentUser {
					getMealsForCurrentUser {
						id 
						name 
						date 
						totalAmount 
						userId 
						description 
						createDate 
						updatedDate 
						userIds 
						products 
					}
				}
				`),
			variables: {
				
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.getMealsForCurrentUser) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.getMealsForCurrentUser, null];
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
						isConfirmed 
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
						admin 
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

async function logout(): Promise<[response | null, any | null]> {
	try {
		const response = await apolloClient.mutate({
			mutation: gql(`
				 mutation logout {
					logout {
						status 
						message 
						success 
					}
				}
				`),
			variables: {
				
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.logout) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.logout, null];
	} catch (error) {
		return [null, error];
	}
}

async function updateUser(user: UpdateUserInput): Promise<[response | null, any | null]> {
	try {
		const response = await apolloClient.mutate({
			mutation: gql(`
				 mutation updateUser($user: UpdateUserInput!) {
					updateUser(user: $user) {
						status 
						message 
						success 
					}
				}
				`),
			variables: {
								user: user
			}
		});
	
		if (response.errors && response.errors.length > 0) {
			return [null, response.errors];
		}
	
		if (!response.data || !response.data.updateUser) {
			return [null, new Error(`Invalid response structure`)];
		}
	
		return [response.data.updateUser, null];
	} catch (error) {
		return [null, error];
	}
}

export const query = {
	getActivities,
	getCurrentUser,
	getDashboard,
	getDebtDetails,
	getDebts,
	getMealById,
	getMeals,
	getMealsForCurrentUser,
	getPayments,
	getUsers
};
export const mutation = {
	createMeal,
	createPayment,
	createPost,
	createUser,
	deleteUser,
	editUser,
	logout,
	updateUser
};


